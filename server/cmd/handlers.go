package main

import (
	"fmt"
	"math/rand"
)

// handler function template
type handlerFn func(*Server, *client, Message)

var registry = map[string]handlerFn{
	"chat":      chat,
	"roll_dice": roll,
}

// dispatcher --> sends message to corresponding function
func dispatch(s *Server, c *client, msg Message) {
	if h, ok := registry[msg.Type]; ok {
		h(s, c, msg)
	} else {
		fmt.Println("unknown message:", msg.Type)
	}
}

// chat function 
func chat(s *Server, sender *client, msg Message) {
	s.broadcast(Message{
		Type:   msg.Type,
		Sender: sender.conn.RemoteAddr().String(),
		Data:   msg.Data,
	})
}

// dice roll function
func roll(s *Server, sender *client, msg Message) {
	d1, d2 := rand.Intn(6)+1, rand.Intn(6)+1

	s.broadcast(Message{
		Type:   "chat",
		Sender: sender.conn.RemoteAddr().String(),
		Data:   fmt.Sprintf("rolled %d & %d", d1, d2),
	})
}
