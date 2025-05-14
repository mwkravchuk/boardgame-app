package main

import (
	"encoding/json"
	"log"
	"sync"
	"github.com/gorilla/websocket"
	"github.com/google/uuid"
	"fmt"
)

// message sharing format
type Message struct {
	Type   string      `json:"type"`
	Sender string      `json:"sender,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

// client struct for individual mutexes
type Client struct {
	conn *websocket.Conn
	id   string
	mu   sync.Mutex
}

// server data structure 
type Server struct {
	clients      map[*Client]bool
	clientsMutex sync.RWMutex
	turnOrder    []string
	currentTurn  int
}

func NewServer() *Server {
	return &Server{clients: make(map[*Client]bool)}
}

// adding and removing clients
func (s *Server) AddClient(ws *websocket.Conn) {
	id := uuid.New().String()
	c := &Client{
		conn: ws,
		id:   id,
	}

	// Server stores map of active clients
	s.clientsMutex.Lock()
	s.clients[c] = true
	s.clientsMutex.Unlock()

	// Tell client their new id
	msg := Message{
		Type: "new_id",
	}
	dispatch(s, c, msg)

	// Add client to the slice of players for this game
	s.turnOrder = append(s.turnOrder, id)

	// Start game when two players join and send message of new turn
	if len(s.turnOrder) == 2 {
		s.currentTurn = 0
		msg = Message{
			Type: "new_turn",
		}
		fmt.Println("game started")
		dispatch(s, c, msg)
	}

	go s.readLoop(c)
}

func (s *Server) removeClient(c *Client) {
	s.clientsMutex.Lock()
	if _, ok := s.clients[c]; ok {
		delete(s.clients, c)
		c.conn.Close()
	}
	s.clientsMutex.Unlock()
}

// read loop
func (s *Server) readLoop(c *Client) {
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

// signal message to just one client
func (s *Server) signal(c *Client, msg Message) {
	// Convert the message data to JSON
	jsonData, err := json.Marshal(msg)
	if err != nil {
		log.Println("Error marshalling message:", err)
		return
	}
	s.clientsMutex.RLock()

	c.mu.Lock()
	if err := c.conn.WriteMessage(websocket.TextMessage, jsonData); err != nil {
		log.Println("write:", err)
		c.mu.Unlock()
		go s.removeClient(c)
	}
	c.mu.Unlock()
	
	s.clientsMutex.RUnlock()
}

// broadcast the message back to all the clients
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
