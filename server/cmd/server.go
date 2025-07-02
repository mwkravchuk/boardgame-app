package main

import (
	"encoding/json"
	"log"
	"sync"
	"github.com/gorilla/websocket"
	"github.com/google/uuid"
)

type GameState struct {
	Players     map[string]*PlayerState `json:"players"`
	TurnOrder   []string 								`json:"turnOrder"`
	CurrentTurn int 										`json:"currentTurn"`
	BoardState  []int 									`json:"boardState"`
	Properties  []Property							`json:"properties"`
}

type Property struct {
	Name 				string `json:"name"`
	Price 	    int    `json:"price"`
	Rent        int		 `json:"rent"`
	OwnerID     string `json:"ownerId"`
	IsProperty  bool   `json:"isProperty"`
	IsOwned     bool   `json:"isOwned"`
	Color       string `json:"color"`
	IsMortgaged bool   `json:"isMortgaged"`
}

type PlayerState struct {
	ID              string 		 `json:"id"`
	DisplayName     string     `json:"displayName"`
	Position        int    		 `json:"position"`
	Money           int    		 `json:"money"`
	InJail          bool   		 `json:"inJail"`
	Color           string 		 `json:"color"`
	PropertiesOwned []int      `json:"properties"` // store indices of owned properties
}

type GameRoom struct {
	Code      string
	Players 	map[*Client]bool
	GameState *GameState
}

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
	Rooms        map[string]*GameRoom // Map from room code to room
	ClientToRoomCode map[*Client]string
}

func NewServer() *Server {
	return &Server{clients: make(map[*Client]bool),
								 Rooms: make(map[string]*GameRoom),
								 ClientToRoomCode: make(map[*Client]string),
								}
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


func (s *Server) broadcastToRoom(room *GameRoom, msg Message) {
		// Convert the message data to JSON
	jsonData, err := json.Marshal(msg)
	if err != nil {
		log.Println("Error marshalling message:", err)
		return
	}

	s.clientsMutex.RLock()
	for c := range room.Players {
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
