package handlers

import (
	"fmt"
	"math/rand"
	"boardgame-app/server/network"
	"boardgame-app/server/shared"
)

func Roll(s *network.Server, sender *shared.Client, msg shared.Message) {
	room, ok := IsInValidRoom(s, sender)
	if !ok {
		return
	}

	player := room.GameState.Players[sender.Id]
	displayName := player.DisplayName

	// Roll the dice
	d1, d2 := rand.Intn(6) + 1, rand.Intn(6) + 1
	totalDice := d1 + d2

	// Update game state
	room.GameState.LastRoll = totalDice
	player.Position = (player.Position + totalDice) % 40
	player.HasRolled = true

	// Send messages based on state of property
	property := &room.GameState.Properties[player.Position]

	if property.IsProperty {
		if !property.IsOwned {
			s.Signal(sender, shared.Message{
				Type: "can_buy_property",
				Data: map[string]interface{}{
					"property": property,
				},
			})
		} else if property.IsOwned && property.OwnerID != player.ID && !property.IsMortgaged {
			rent := CalculateRent(room.GameState.Players[property.OwnerID], player.Position, room.GameState)
			s.Signal(sender, shared.Message{
				Type: "owe_rent",
				Data: map[string]interface{}{
					"property": property,
					"rent": rent,
					"displayName": room.GameState.Players[property.OwnerID].DisplayName,
				},
			})
		}
	}

	s.BroadcastToRoom(room, shared.Message{
		Type:   "dice_rolled",
		Sender: displayName,
		Data:   []int{d1, d2},
	})

	s.BroadcastToRoom(room, shared.Message{
		Type:   "console",
		Sender: displayName,
		Data:   fmt.Sprintf("rolled %d & %d", d1, d2),
	})

	s.BroadcastToRoom(room, shared.Message{
		Type:   "game_state",
		Sender: sender.Conn.RemoteAddr().String(),
		Data:   room.GameState,
	})
}

func NewTurn(s *network.Server, sender *shared.Client, msg shared.Message) {
	// Find the correct room
	room, ok := IsInValidRoom(s, sender)
	if !ok {
		return
	}

	room.GameState.CurrentTurn = (room.GameState.CurrentTurn + 1) % len(room.GameState.Players)
	player := room.GameState.Players[sender.Id]
	player.HasRolled = false

	s.BroadcastToRoom(room, shared.Message{
		Type:   "game_state",
		Sender: sender.Conn.RemoteAddr().String(),
		Data:   room.GameState,
	})
}