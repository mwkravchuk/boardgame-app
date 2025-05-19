package main

import (
	"fmt"
	"math/rand"
	"time"
)

// handler function template
type handlerFn func(*Server, *Client, Message)

var registry = map[string]handlerFn{
	"start_game":  startGame,
	"join_room":   joinRoom,
	"create_room": createRoom,
	"new_turn":    newTurn,
	"new_id":      newId,
	"chat":        chat,
	"roll_dice":   roll,
}

// dispatcher --> sends message to corresponding function
func dispatch(s *Server, c *Client, msg Message) {
	if h, ok := registry[msg.Type]; ok {
		h(s, c, msg)
	} else {
		fmt.Println("unknown message:", msg.Type)
	}
}

func startGame(s *Server, sender *Client, msg Message) {

	roomCode, ok := s.ClientToRoomCode[sender]
	if !ok {
		fmt.Println("Client not in a room")
		return
	}

	room, ok := s.Rooms[roomCode]
	if !ok {
		fmt.Println("Room not found: ", roomCode)
		return
	}

	// Need at least 2 players
	if len(room.Players) < 2 {
		fmt.Println("Not enough players to start")
		return
	}

	// This is the "party leader" for now
	if room.TurnOrder[0] != sender.id {
		fmt.Println("Only room creator can start game")
		return
	}

	// Tell everyone in the room that the game started
	s.broadcastToRoom(room, Message{
		Type:   "game_started",
		Sender: sender.conn.RemoteAddr().String(),
	})
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

func createRoom(s *Server, sender *Client, msg Message) {
	// Create game room
	code := generateRoomCode()
	room := &GameRoom{
		Code:        code,
		Players:     make(map[*Client]bool),
		TurnOrder:   []string{},
		CurrentTurn: -1,
	}

	// Add the sender to the room
	room.Players[sender] = true
	room.TurnOrder = append(room.TurnOrder, sender.id)
	s.ClientToRoomCode[sender] = code

	// Add the room to the server
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

		if room, ok := s.Rooms[code]; ok {
			// Room exists. Add them to it and let them know
			room.Players[sender] = true
			room.TurnOrder = append(room.TurnOrder, sender.id)
			s.ClientToRoomCode[sender] = code

			fmt.Println("client joined room: ", room)

			s.signal(sender, Message{
				Type:   "room_joined",
				Sender: sender.conn.RemoteAddr().String(),
				Data:   code,
			})
		} else {
			// Room doesn't exist. Do nothing.
		}
}

func newTurn(s *Server, sender *Client, msg Message) {

	currentTurnId := s.turnOrder[s.currentTurn % len(s.turnOrder)]

	s.broadcast(Message{
		Type:   "new_turn",
		Sender: sender.conn.RemoteAddr().String(),
		Data:   currentTurnId,
	})

	// logic to reset "hasRolled" dice on client side
	s.signal(sender, Message{
		Type: "reset_roll_button",
		Data: false,
	})

	s.currentTurn += 1

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
	s.broadcast(Message{
		Type:   "chat",
		Sender: sender.conn.RemoteAddr().String(),
		Data:   msg.Data,
	})
}

// dice roll function
func roll(s *Server, sender *Client, msg Message) {
	d1, d2 := rand.Intn(6) + 1, rand.Intn(6) + 1

	// confirm to sender that dice has been rolled
	s.signal(sender, Message{
		Type: "roll_dice",
		Data: true,
	})

	s.broadcast(Message{
		Type:   "chat",
		Sender: sender.conn.RemoteAddr().String(),
		Data:   fmt.Sprintf("rolled %d & %d", d1, d2),
	})
}
