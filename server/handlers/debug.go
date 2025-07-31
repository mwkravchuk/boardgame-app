package handlers

import (
	"fmt"
	"boardgame-app/server/network"
	"boardgame-app/server/shared"
)

func DebugGiveProperty(s *network.Server, sender *shared.Client, msg shared.Message) {
	if !s.DebugMode {
		return // ignore in production
	}

	room, ok := IsInValidRoom(s, sender)
	if !ok {
		return
	}

	propertyIdx := int(msg.Data.(float64))

	// Look up player
	player, exists := room.GameState.Players[sender.Id]
	if !exists {
		fmt.Println("Player not found")
		return
	}

	// Give property if not already owned
	alreadyOwned := false
	for _, p := range player.PropertiesOwned {
		if p == propertyIdx {
			alreadyOwned = true
			break
		}
	}

	if !alreadyOwned {
		player.PropertiesOwned = append(player.PropertiesOwned, propertyIdx)
	}

	// Optional: broadcast updated state
	s.BroadcastToRoom(room, shared.Message{
		Type:   "game_state",
		Sender: sender.Conn.RemoteAddr().String(),
		Data:   room.GameState,
	})
}