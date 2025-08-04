package handlers

import (
	"fmt"
	"boardgame-app/server/network"
	"boardgame-app/server/shared"
)

func StartGame(s *network.Server, sender *shared.Client, msg shared.Message) {
	room, ok := IsInValidRoom(s, sender)
	if !ok {
		return
	}

	// ERROR: Need at least 2 players. Only game creator can start
	if len(room.PlayerIDs) < 2 || room.GameState.TurnOrder[0] != sender.Id {
		s.Signal(sender, shared.Message{
			Type:   "game_started_fail",
			Sender: sender.Conn.RemoteAddr().String(),
		})
		return
	}

	// Indicate game start
	s.BroadcastToRoom(room, shared.Message{
		Type:   "game_started",
		Sender: sender.Conn.RemoteAddr().String(),
	})
}

func BuyProperty(s *network.Server, sender *shared.Client, msg shared.Message) {
	fmt.Println("Buying property")
	room, ok := IsInValidRoom(s, sender)
	if !ok {
		return
	}

	playerId := sender.Id
	player := room.GameState.Players[sender.Id]
	propertyIndex := player.Position
	property := &room.GameState.Properties[propertyIndex]

	// update property
	property.IsOwned = true
	property.OwnerID = playerId
	player.Money -= property.Price
	player.PropertiesOwned = append(player.PropertiesOwned, propertyIndex)

	s.BroadcastToRoom(room, shared.Message{
		Type:   "game_state",
		Sender: sender.Conn.RemoteAddr().String(),
		Data:   room.GameState,
	})
}

func AuctionProperty(s *network.Server, sender *shared.Client, msg shared.Message) {
	fmt.Println("Auctioning property", msg.Data)
}

func PayRent(s *network.Server, sender *shared.Client, msg shared.Message) {
	fmt.Println("Paying rent", msg.Data)
	room, ok := IsInValidRoom(s, sender)
	if !ok {
		return
	}

	player := room.GameState.Players[sender.Id]
	propertyIndex := player.Position
	property := &room.GameState.Properties[propertyIndex]

	if property.OwnerID == "" || property.OwnerID == sender.Id {
		return
	}

	propertyOwner := room.GameState.Players[property.OwnerID]
	rent := CalculateRent(propertyOwner, propertyIndex, room.GameState)

	// Send the money
	player.Money -= rent
	propertyOwner.Money += rent

	s.BroadcastToRoom(room, shared.Message{
		Type:   "game_state",
		Sender: sender.Conn.RemoteAddr().String(),
		Data:   room.GameState,
	})
}

func Bankrupt(s *network.Server, sender *shared.Client, msg shared.Message) {
	fmt.Println("Bankrupting")
}