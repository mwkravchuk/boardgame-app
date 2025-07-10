package handlers

import (
	"fmt"
	"boardgame-app/server/network"
	"boardgame-app/server/shared"
)

func NewId(s *network.Server, sender *shared.Client, msg shared.Message) {
	s.Signal(sender, shared.Message{
		Type:   "new_id",
		Sender: sender.Conn.RemoteAddr().String(),
		Data:   sender.Id,
	})
	fmt.Println("New user created with id: ", sender.Id)
}

func CreateRoom(s *network.Server, sender *shared.Client, msg shared.Message) {
	// Create game room
	code := GenerateRoomCode()

	data, ok := msg.Data.(map[string]interface{})
	if !ok {
		fmt.Println("Invalid message data format")
		return
	}

	displayNameRaw, ok := data["displayName"]
	if !ok {
		fmt.Println("displayName missing")
		return
	}

	displayName, ok := displayNameRaw.(string)
	if !ok {
		fmt.Println("displayName is not a string")
		return
	}

	// Initiate game/player state information
	playerId := sender.Id
	initialPlayerState := &shared.PlayerState{
		ID:          playerId,
		DisplayName: displayName,
		Position:    0,
		Money:       1500,
		InJail:      false,
	}

	initialGameState := &shared.GameState{
		Players: map[string]*shared.PlayerState{
			playerId: initialPlayerState,
		},
		TurnOrder: []string{playerId},
		CurrentTurn: 0,
		BoardState: make([]int, 40),
		Properties: InitializeProperties(),
	}

	// Create room
	room := &shared.GameRoom{
		Code:      code,
		PlayerIDs:   make(map[string]bool),
		GameState: initialGameState,
	}

	// Consider sender as first player to this game. Add to appropriate maps
	room.PlayerIDs[sender.Id] = true
	s.ClientToRoomCode[sender.Id] = code

	// Add the room to the server's list of rooms
	s.Rooms[code] = room

	// Sender joins the room after it was created. Let them know!
	s.Signal(sender, shared.Message{
		Type:   "room_joined",
		Sender: sender.Conn.RemoteAddr().String(),
		Data:   code,
	})
}

func JoinRoom(s *network.Server, sender *shared.Client, msg shared.Message) {

	fmt.Println("Joined room")
	
	data, ok := msg.Data.(map[string]interface{})
	if !ok {
		fmt.Println("Invalid message data format")
		return
	}

	codeRaw, ok := data["code"]
	if !ok {
		fmt.Println("join code missing")
		return
	}

	code, ok := codeRaw.(string)
	if !ok {
		fmt.Println("join code is not a string")
		return
	}

	displayNameRaw, ok := data["displayName"]
	if !ok {
		fmt.Println("displayName missing")
		return
	}

	displayName, ok := displayNameRaw.(string)
	if !ok {
		fmt.Println("displayName is not a string")
		return
	}

	playerId := sender.Id

	if room, ok := s.Rooms[code]; ok {
		// Room exists. Add them to it and let them know
		room.PlayerIDs[playerId] = true

		// Initialize new player
		newPlayer := &shared.PlayerState{
			ID:          playerId,
			DisplayName: displayName,
			Position:    0,
			Money:       1500,
			InJail:      false,
		}
		room.GameState.Players[playerId] = newPlayer

		// Add to appropriate maps
		room.GameState.TurnOrder = append(room.GameState.TurnOrder, playerId)
		s.ClientToRoomCode[playerId] = code

		s.Signal(sender, shared.Message{
			Type:   "room_joined",
			Sender: sender.Conn.RemoteAddr().String(),
			Data:   code,
		})
	}
}

func ColorSelected(s *network.Server, sender *shared.Client, msg shared.Message) {
	room, ok := IsInValidRoom(s, sender)
	if !ok {
		return
	}
	color := msg.Data.(string)
	room.GameState.Players[sender.Id].Color = color
	fmt.Println("color:", room.GameState.Players[sender.Id].Color)
}

