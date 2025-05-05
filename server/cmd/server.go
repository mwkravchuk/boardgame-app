package main

import (
	"encoding/json"
	"log"
	"sync"
	"github.com/gorilla/websocket"
)

// message sharing format
type Message struct {
	Type   string      `json:"type"`
	Sender string      `json:"sender,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

// client struct for individual mutexes
type client struct {
	conn *websocket.Conn
	mu   sync.Mutex
}

// server data structure 
type Server struct {
	clients      map[*client]bool
	clientsMutex sync.RWMutex
}

func NewServer() *Server {
	return &Server{clients: make(map[*client]bool)}
}

// adding and removing clients
func (s *Server) AddClient(ws *websocket.Conn) {
	c := &client{conn: ws}
	s.clientsMutex.Lock()
	s.clients[c] = true
	s.clientsMutex.Unlock()

	go s.readLoop(c)
}

func (s *Server) removeClient(c *client) {
	s.clientsMutex.Lock()
	if _, ok := s.clients[c]; ok {
		delete(s.clients, c)
		c.conn.Close()
	}
	s.clientsMutex.Unlock()
}

// read loop
func (s *Server) readLoop(c *client) {
	defer s.removeClient(c)

	for {
		var msg Message
		if err := c.conn.ReadJSON(&msg); err != nil {
			log.Println("read:", err)
			return
		}
		dispatch(s, c, msg) // find corresponding function in handlers.go
	}
}

// broadcast the message back to the client
func (s *Server) broadcast(msg Message) {


	// Convert the message data to JSON
	jsonData, err := json.Marshal(msg)
	if err != nil {
		log.Println("Error marshalling message:", err)
		return
	}

	s.clientsMutex.RLock()
	for c := range s.clients {
		c.mu.Lock()
		if err := c.conn.WriteMessage(websocket.TextMessage, jsonData); err != nil {
			log.Println("write:", err)
			c.mu.Unlock()
			go s.removeClient(c)
			continue
		}
		c.mu.Unlock()
	}
	s.clientsMutex.RUnlock()
}
