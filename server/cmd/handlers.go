package main

import (
	"fmt"
	"math/rand"
)

// handler function template
type handlerFn func(*Server, *Client, Message)

var registry = map[string]handlerFn{
	"new_turn":  newTurn,
	"new_id":    newId,
	"chat":      chat,
	"roll_dice": roll,
}

// dispatcher --> sends message to corresponding function
func dispatch(s *Server, c *Client, msg Message) {
	if h, ok := registry[msg.Type]; ok {
		h(s, c, msg)
	} else {
		fmt.Println("unknown message:", msg.Type)
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
