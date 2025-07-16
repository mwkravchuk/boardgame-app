package handlers

import (
	"fmt"
	"math/rand"
	"boardgame-app/server/network"
	"boardgame-app/server/shared"
)

func Roll(s *network.Server, sender *shared.Client, msg shared.Message) {
	room, ok := IsInValidRoom(s, sender)

	displayName := room.GameState.Players[sender.Id].DisplayName
	if !ok {
		return
	}

	d1, d2 := rand.Intn(6) + 1, rand.Intn(6) + 1
	totalDice := d1 + d2
	playerId := sender.Id
	room.GameState.Players[playerId].Position = (room.GameState.Players[playerId].Position + totalDice) % 40

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

	room.GameState.CurrentTurn += 1

	s.BroadcastToRoom(room, shared.Message{
		Type:   "game_state",
		Sender: sender.Conn.RemoteAddr().String(),
		Data:   room.GameState,
	})

	// logic to reset "hasRolled" dice on client side
	s.Signal(sender, shared.Message{
		Type: "reset_roll_button",
		Data: false,
	})
}