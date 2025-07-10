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
	// Need at least 2 players
	if len(room.PlayerIDs) < 2 {
		fmt.Println("Not enough players to start")
		return
	}

	// This is the "party leader" for now
	if room.GameState.TurnOrder[0] != sender.Id {
		fmt.Println("Only room creator can start game")
		return
	}

	// Tell everyone in the room that the game started
	s.BroadcastToRoom(room, shared.Message{
		Type:   "game_started",
		Sender: sender.Conn.RemoteAddr().String(),
	})
	
	// Send a new_turn message to the party owner to start game.
	NewTurn(s, sender, msg)

	s.BroadcastToRoom(room, shared.Message{
		Type:   "game_state",
		Sender: sender.Conn.RemoteAddr().String(),
		Data:   room.GameState,
	})
}

func BuyProperty(s *network.Server, sender *shared.Client, msg shared.Message) {
	fmt.Println("buy property data: ", msg.Data)
	room, ok := IsInValidRoom(s, sender)
	if !ok {
		return
	}

	playerId := sender.Id
	player := room.GameState.Players[playerId]
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
	fmt.Println("auction property data: ", msg.Data)
}

func PayRent(s *network.Server, sender *shared.Client, msg shared.Message) {
	fmt.Println("pay rent data: ", msg.Data)
	room, ok := IsInValidRoom(s, sender)
	if !ok {
		return
	}

	playerId := sender.Id
	player := room.GameState.Players[playerId]
	propertyIndex := player.Position
	property := &room.GameState.Properties[propertyIndex]
	propertyOwner := room.GameState.Players[property.OwnerID]

	// send money to deserved player
	player.Money -= property.Rent
	propertyOwner.Money += property.Rent

	s.BroadcastToRoom(room, shared.Message{
		Type:   "game_state",
		Sender: sender.Conn.RemoteAddr().String(),
		Data:   room.GameState,
	})
}