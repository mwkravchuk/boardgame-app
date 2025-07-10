package network

import (
	"encoding/json"
	"log"
	"sync"
	"boardgame-app/server/shared"
	"github.com/gorilla/websocket"
)

type Server struct {
	ConnectedIds     map[string]bool
	ClientsById      map[string]*shared.Client
	ClientToRoomCode map[string]string    			 // clientId -> roomCode
	Rooms            map[string]*shared.GameRoom // roomCode -> room
	ClientsMutex     sync.RWMutex
}

func NewServer() *Server {
	return &Server{ConnectedIds:     make(map[string]bool),
								 ClientsById:      make(map[string]*shared.Client),
								 ClientToRoomCode: make(map[string]string),
								 Rooms:            make(map[string]*shared.GameRoom),
								}
}

// Share this message to ONE client
func (s *Server) Signal(client *shared.Client, msg shared.Message) {
	jsonData, err := json.Marshal(msg)
	if err != nil {
		log.Println("Error marshalling message:", err)
		return
	}

	s.ClientsMutex.RLock()
	defer s.ClientsMutex.RUnlock()

	client.Mu.Lock()
	err = client.Conn.WriteMessage(websocket.TextMessage, jsonData)
	client.Mu.Unlock()

	if err != nil {
		log.Printf("Write to client %s failed: %v\n", client.Id, err)
		go s.RemoveClient(client)
	}
}

// Share this message to ALL clients
func (s *Server) Broadcast(msg shared.Message) {
	jsonData, err := json.Marshal(msg)
	if err != nil {
		log.Println("Error marshalling message:", err)
		return
	}

	s.ClientsMutex.RLock()
	defer s.ClientsMutex.RUnlock()

	for _, client := range s.ClientsById {
		client.Mu.Lock()
		err = client.Conn.WriteMessage(websocket.TextMessage, jsonData)
		client.Mu.Unlock()

		if err != nil {
			log.Printf("Write to client %s failed: %v\n", client.Id, err)
			go s.RemoveClient(client)
		}
	}
}

// Share this message to ALL clients WITHIN A ROOM
func (s *Server) BroadcastToRoom(room *shared.GameRoom, msg shared.Message) {
	jsonData, err := json.Marshal(msg)
	if err != nil {
		log.Println("Error marshalling message:", err)
		return
	}

	s.ClientsMutex.RLock()
	defer s.ClientsMutex.RUnlock()
	for clientId := range room.PlayerIDs {
		client, ok := s.ClientsById[clientId]
		if !ok {
			log.Printf("Client with ID %s not found\n", clientId)
			continue
		}

		client.Mu.Lock()
		err := client.Conn.WriteMessage(websocket.TextMessage, jsonData)
		client.Mu.Unlock()

		if err != nil {
			log.Printf("Write to client %s failed: %v\n", client.Id, err)
			go s.RemoveClient(client)
		}
	}
}

