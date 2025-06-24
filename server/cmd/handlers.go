package main

import (
	"fmt"
	"math/rand"
	"time"
)

// handler function template
type handlerFn func(*Server, *Client, Message)

var registry = map[string]handlerFn{
	"start_game":     startGame,
	"join_room":      joinRoom,
	"create_room":    createRoom,
	"new_turn":       newTurn,
	"new_id":         newId,
	"chat":           chat,
	"console":        console,
	"roll_dice":      roll,
	"color_selected": colorSelected,
}

func colorSelected(s *Server, sender *Client, msg Message) {
	room, ok := isInValidRoom(s, sender)
	if ok {
		color := msg.Data.(string)
		room.GameState.Players[sender.id].Color = color
		fmt.Println("color:", room.GameState.Players[sender.id].Color)
	}
}

// dispatcher --> sends message to corresponding function
func dispatch(s *Server, c *Client, msg Message) {
	if h, ok := registry[msg.Type]; ok {
		h(s, c, msg)
	} else {
		fmt.Println("unknown message:", msg.Type)
	}
}

func isInValidRoom(s *Server, sender *Client) (*GameRoom, bool) {
	roomCode, ok := s.ClientToRoomCode[sender]
	if !ok {
		fmt.Println("Client not in a room")
		return nil, false
	}

	room, ok := s.Rooms[roomCode]
	if !ok {
		fmt.Println("Room not found: ", roomCode)
		return nil, false
	}

	return room, true
}

func startGame(s *Server, sender *Client, msg Message) {

	room, ok := isInValidRoom(s, sender)
	if ok {
		// Need at least 2 players
		if len(room.Players) < 2 {
			fmt.Println("Not enough players to start")
			return
		}

		// This is the "party leader" for now
		if room.GameState.TurnOrder[0] != sender.id {
			fmt.Println("Only room creator can start game")
			return
		}

		// Tell everyone in the room that the game started
		s.broadcastToRoom(room, Message{
			Type:   "game_started",
			Sender: sender.conn.RemoteAddr().String(),
		})
		
		// Send a new_turn message to the party owner to start game.
		newTurn(s, sender, msg)

		s.broadcastToRoom(room, Message{
			Type:   "game_state",
			Sender: sender.conn.RemoteAddr().String(),
			Data:   room.GameState,
		})
	}
}

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const codeLength = 4

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateRoomCode() string {
	b := make([]byte, codeLength)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func initializeProperties() []Property {
	properties := make([]Property, 40)
	properties[1] = Property{
		Name: "Mediterranean Avenue",
		Color: "brown",
		Price: 60,
		Rent: 2,
		OwnerID: "",
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[3] = Property{
		Name: "Baltic Avenue",
		Color: "brown",
		Price: 60,
		Rent: 4,
		OwnerID: "",
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[5] = Property{
		Name: "Reading Railroad",
		Color: "black",
		Price: 200,
		Rent: 25,
		OwnerID: "",
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[6] = Property{
		Name: "Oriental Avenue",
		Color: "light blue",
		Price: 100,
		Rent: 6,
		OwnerID: "",
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[8] = Property{
		Name: "Vermont Avenue",
		Color: "light blue",
		Price: 100,
		Rent: 6,
		OwnerID: "",
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[9] = Property{
		Name: "Connecticut Avenue",
		Color: "light blue",
		Price: 120,
		Rent: 8,
		OwnerID: "",
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[11] = Property{
		Name: "St. Charles Place",
		Color: "pink",
		Price: 140,
		Rent: 10,
		OwnerID: "",
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[12] = Property{
		Name: "Electric Company",
		Color: "white",
		Price: 150,
		Rent: 0,
		OwnerID: "",
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[13] = Property{
		Name: "States Avenue",
		Color: "pink",
		Price: 140,
		Rent: 10,
		OwnerID: "",
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[14] = Property{
		Name: "Virginia Avenue",
		Color: "pink",
		Price: 160,
		Rent: 12,
		OwnerID: "",
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[15] = Property{
		Name: "Pennsylvania Railroad",
		Color: "black",
		Price: 200,
		Rent: 25,
		OwnerID: "",
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[16] = Property{
		Name: "St. James Place",
		Color: "orange",
		Price: 180,
		Rent: 14,
		OwnerID: "",
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[18] = Property{
		Name: "Tennessee Avenue",
		Color: "orange",
		Price: 180,
		Rent: 14,
		OwnerID: "",
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[19] = Property{
		Name: "New York Avenue",
		Color: "orange",
		Price: 200,
		Rent: 16,
		OwnerID: "",
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[21] = Property{
		Name: "Kentucky Avenue",
		Color: "red",
		Price: 220,
		Rent: 18,
		OwnerID: "",
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[23] = Property{
		Name: "Indiana Avenue",
		Color: "red",
		Price: 220,
		Rent: 18,
		OwnerID: "",
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[24] = Property{
		Name: "Illinois Avenue",
		Color: "red",
		Price: 240,
		Rent: 20,
		OwnerID: "",
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[25] = Property{
		Name: "B & O Railroad",
		Color: "black",
		Price: 200,
		Rent: 25,
		OwnerID: "",
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[26] = Property{
		Name: "Atlantic Avenue",
		Color: "yellow",
		Price: 260,
		Rent: 22,
		OwnerID: "",
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[27] = Property{
		Name: "Ventnor Avenue",
		Color: "yellow",
		Price: 260,
		Rent: 22,
		OwnerID: "",
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[28] = Property{
		Name: "Water Works",
		Color: "white",
		Price: 150,
		Rent: 0,
		OwnerID: "",
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[29] = Property{
		Name: "Marvin Gardens",
		Color: "yellow",
		Price: 280,
		Rent: 24,
		OwnerID: "",
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[31] = Property{
		Name: "Pacific Avenue",
		Color: "green",
		Price: 300,
		Rent: 26,
		OwnerID: "",
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[32] = Property{
		Name: "North Carolina Avenue",
		Color: "green",
		Price: 300,
		Rent: 26,
		OwnerID: "",
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[34] = Property{
		Name: "Pennsylvania Avenue",
		Color: "green",
		Price: 320,
		Rent: 28,
		OwnerID: "",
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[35] = Property{
		Name: "Short Line",
		Color: "black",
		Price: 200,
		Rent: 25,
		OwnerID: "",
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[37] = Property{
		Name: "Park Place",
		Color: "blue",
		Price: 350,
		Rent: 35,
		OwnerID: "",
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[39] = Property{
		Name: "Boardwalk",
		Color: "blue",
		Price: 400,
		Rent: 50,
		OwnerID: "",
		IsOwned: false,
		IsMortgaged: false,
	}

	return properties
}

func createRoom(s *Server, sender *Client, msg Message) {
	// Create game room
	code := generateRoomCode()

	// Initiate game/player state information
	playerId := sender.id
	initialPlayerState := &PlayerState{
		ID:       playerId,
		Position: 0,
		Money:    1500,
		InJail:   false,
	}

	initialGameState := &GameState{
		Players: map[string]*PlayerState{
			playerId: initialPlayerState,
		},
		TurnOrder: []string{playerId},
		CurrentTurn: 0,
		BoardState: make([]int, 40),
		Properties: initializeProperties(),
	}

	// Create room
	room := &GameRoom{
		Code:      code,
		Players:   make(map[*Client]bool),
		GameState: initialGameState,
	}

	// Consider sender as first player to this game. Add to appropriate maps
	room.Players[sender] = true
	s.ClientToRoomCode[sender] = code

	// Add the room to the server's list of rooms
	s.Rooms[code] = room

	// Sender joins the room after it was created. Let them know!
	s.signal(sender, Message{
		Type:   "room_joined",
		Sender: sender.conn.RemoteAddr().String(),
		Data:   code,
	})
}

func joinRoom(s *Server, sender *Client, msg Message) {

	fmt.Println("join_room message data: ", msg.Data)
	code := msg.Data.(string)
	playerId := sender.id

	if room, ok := s.Rooms[code]; ok {
		// Room exists. Add them to it and let them know
		room.Players[sender] = true

		// Initialize new player
		newPlayer := &PlayerState{
			ID:       playerId,
			Position: 0,
			Money:    1500,
			InJail:   false,
		}
		room.GameState.Players[playerId] = newPlayer

		// Add to appropriate maps
		room.GameState.TurnOrder = append(room.GameState.TurnOrder, playerId)
		s.ClientToRoomCode[sender] = code

		fmt.Println("client joined room: ", room)
		fmt.Println("gamestate turnorder: ", room.GameState.TurnOrder)
		fmt.Println("gamestate players: ", room.GameState.Players)

		s.signal(sender, Message{
			Type:   "room_joined",
			Sender: sender.conn.RemoteAddr().String(),
			Data:   code,
		})
	}
}

func newTurn(s *Server, sender *Client, msg Message) {

	// Find the correct room
	room, ok := isInValidRoom(s, sender)
	if ok {
		// Update the current turn for this room
		currentTurnId := room.GameState.TurnOrder[room.GameState.CurrentTurn % len(room.GameState.TurnOrder)]
		fmt.Println("New turn message: ", currentTurnId)

		// Tell players that current turn has updated
		s.broadcastToRoom(room, Message{
			Type:   "new_turn",
			Sender: sender.conn.RemoteAddr().String(),
			Data:   currentTurnId,
		})

		// logic to reset "hasRolled" dice on client side
		s.signal(sender, Message{
			Type: "reset_roll_button",
			Data: false,
		})

		room.GameState.CurrentTurn += 1
	}
}

func newId(s *Server, sender *Client, msg Message) {
	s.signal(sender, Message{
		Type:   "new_id",
		Sender: sender.conn.RemoteAddr().String(),
		Data:   sender.id,
	})
	fmt.Println("New user created with id: ", sender.id)
}

// chat function 
func chat(s *Server, sender *Client, msg Message) {
	room, ok := isInValidRoom(s, sender)
	if ok {
		s.broadcastToRoom(room, Message{
			Type:   "chat",
			Sender: sender.conn.RemoteAddr().String(),
			Data:   msg.Data,
		})
	}
}

func console(s *Server, sender *Client, msg Message) {
	room, ok := isInValidRoom(s, sender)
	if ok {
		s.broadcastToRoom(room, Message{
			Type:   "console",
			Sender: sender.conn.RemoteAddr().String(),
			Data:   msg.Data,
		})	
	}
}

// dice roll function
func roll(s *Server, sender *Client, msg Message) {
	room, ok := isInValidRoom(s, sender)
	if ok {
		d1, d2 := rand.Intn(6) + 1, rand.Intn(6) + 1
		totalDice := d1 + d2
		playerId := sender.id
		room.GameState.Players[playerId].Position = (room.GameState.Players[playerId].Position + totalDice) % 40

		// confirm to sender that dice has been rolled
		s.signal(sender, Message{
			Type: "roll_dice",
			Data: true,
		})

		s.broadcastToRoom(room, Message{
			Type:   "console",
			Sender: sender.conn.RemoteAddr().String(),
			Data:   fmt.Sprintf("rolled %d & %d", d1, d2),
		})

		s.broadcastToRoom(room, Message{
			Type:   "game_state",
			Sender: sender.conn.RemoteAddr().String(),
			Data:   room.GameState,
		})
	}
}
