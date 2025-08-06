package handlers

import (
	"fmt"
	"encoding/json"
	"boardgame-app/server/network"
	"boardgame-app/server/shared"
)

func ProposeTrade(s *network.Server, sender *shared.Client, msg shared.Message) {

	fmt.Println("proposing trade")

	room, ok := IsInValidRoom(s, sender)
	if !ok {
		return
	}

	// Decode received message into struct
	bytes, err := json.Marshal(msg.Data)
	if err != nil {
		fmt.Println("Failed to marshal trade data:", err)
		return
	}

	var payload shared.ProposeTradePayload
	if err := json.Unmarshal(bytes, &payload); err != nil {
		fmt.Println("Invalid trade payload:", err)
		return
	}

	fmt.Println("payload unmarshaled", payload)

	// Validate required fields
	if payload.TargetID == "" {
		fmt.Println("TargetID is required")
		return
	}

	game := room.GameState
	if _, ok := game.Players[payload.TargetID]; !ok {
		fmt.Println("Target player does not exist")
		return
	}

	// Validate money (does initiator have enough?)
	senderPlayer := game.Players[sender.Id]
	if senderPlayer.Money < payload.MyOfferMoney {
		fmt.Println("Not enough money to offer")
		return
	}

	// (TODO: Validate property ownership when you add properties)

	// Store the trade as pending
	trade := &shared.TradeOffer{
		FromPlayerID: sender.Id,
		ToPlayerID:   payload.TargetID,
		OfferMoney:   payload.MyOfferMoney,
		RequestMoney: payload.TheirOfferMoney,
		OfferProps:   payload.MyOfferProps,
		RequestProps: payload.TheirOfferProps,
		Status:       "pending",
	}
	game.CurrentTrade = trade

	// Send prompt to recipient
	targetPlayer := room.Clients[payload.TargetID]
	s.Signal(targetPlayer, shared.Message{
		Type: "trade_offered",
		Data: map[string]interface{}{
			"fromPlayer":   senderPlayer.DisplayName,
			"offerMoney":   payload.MyOfferMoney,
			"requestMoney": payload.TheirOfferMoney,
			"offerProps":   payload.MyOfferProps,
			"requestProps": payload.TheirOfferProps,
		},
	})
}

func RespondToTrade(s *network.Server, sender *shared.Client, msg shared.Message) {
	// 1. Parse response
	response, ok := msg.Data.(string)
	if !ok {
		fmt.Println("Invalid trade response")
		return
	}

	// 2. Retrieve room and game
	room, ok := IsInValidRoom(s, sender)
	if !ok {
		fmt.Println("Player not in a room")
		return
	}
	game := room.GameState
	trade := game.CurrentTrade
	if trade == nil {
		fmt.Println("No active trade to respond to")
		return
	}

	// Validate players
	fromPlayer := game.Players[trade.FromPlayerID]
	toPlayer := game.Players[trade.ToPlayerID]
	if fromPlayer == nil || toPlayer == nil {
		fmt.Println("One of the trade players no longer exists")
		return
	}

	// 3. Validate responder
  if sender.Id != trade.ToPlayerID {
    fmt.Println("Only trade recipient can respond")
    return
  }

	// 4. Handle reject
	if response == "reject" {
		trade.Status = "rejected"
		game.CurrentTrade = nil
		s.Signal(room.Clients[fromPlayer.ID], shared.Message{
			Type: "trade_rejected",
			Sender: sender.Conn.RemoteAddr().String(),
		})

		// Broadcast updated state (clears trade dialog)
		s.BroadcastToRoom(room, shared.Message{
			Type:   "game_state",
			Sender: sender.Conn.RemoteAddr().String(),
			Data:   room.GameState,
		})
		return
	}

	// 5. Handle accept
	if response == "accept" {
		// Validate money
		if fromPlayer.Money < trade.OfferMoney || toPlayer.Money < trade.RequestMoney {
			fmt.Println("One of the players no longer has required money")
			return
		}

		trade.Status = "accepted"
		game.CurrentTrade = nil

		s.Signal(room.Clients[fromPlayer.ID], shared.Message{
			Type: "trade_rejected",
			Sender: sender.Conn.RemoteAddr().String(),
		})

		// Transfer money
		fromPlayer.Money -= trade.OfferMoney
		toPlayer.Money += trade.OfferMoney

		toPlayer.Money -= trade.RequestMoney
		fromPlayer.Money += trade.RequestMoney

		// Transfer properties (if implemented)
		transferProperties(trade.OfferProps, fromPlayer, toPlayer)
		transferProperties(trade.RequestProps, toPlayer, fromPlayer)

		// Broadcast updated state
		s.BroadcastToRoom(room, shared.Message{
			Type:   "game_state",
			Sender: sender.Conn.RemoteAddr().String(),
			Data:   room.GameState,
		})
	}
}