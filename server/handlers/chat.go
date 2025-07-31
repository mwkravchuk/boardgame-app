package handlers

import (
	"boardgame-app/server/network"
	"boardgame-app/server/shared"
)

func Chat(s *network.Server, sender *shared.Client, msg shared.Message) {
	room, ok := IsInValidRoom(s, sender)
	displayName := room.GameState.Players[sender.Id].DisplayName
	if !ok {
		return
	}
	s.BroadcastToRoom(room, shared.Message{
		Type:   "chat",
		Sender: displayName,
		Data:   msg.Data,
	})
}

func Console(s *network.Server, sender *shared.Client, msg shared.Message) {
	room, ok := IsInValidRoom(s, sender)
	if !ok {
		return
	}

	displayName := room.GameState.Players[sender.Id].DisplayName

	s.BroadcastToRoom(room, shared.Message{
		Type:   "console",
		Sender: displayName,
		Data:   msg.Data,
	})	
}