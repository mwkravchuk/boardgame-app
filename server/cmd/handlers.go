package main

import (
	"fmt"
	"math/rand"
)

// handler function template
type handlerFn func(*Server, *Client, Message)

var registry = map[string]handlerFn{
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

func newId(s *Server, sender *Client, msg Message) {
	s.broadcast(Message{
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

	s.broadcast(Message{
		Type:   "chat",
		Sender: sender.conn.RemoteAddr().String(),
		Data:   fmt.Sprintf("rolled %d & %d", d1, d2),
	})
}
