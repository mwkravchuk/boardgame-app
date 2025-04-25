package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/websocket"
	"sync"
	"math/rand"
)

type Server struct {
	clients map[*websocket.Conn]bool // store active connections
	clientsMutex sync.Mutex
}

func NewServer() *Server {
	return &Server {
		clients: make(map[*websocket.Conn]bool),
	}
}


type Message struct {
	Type string      `json:"type"`
	Sender string    `json:"sender"`
	Data interface{} `json:"data"`
}

var messageHandlers = map[string]func(*Server, *websocket.Conn, interface{}){
	"chat":       handleChatMessage,
	"roll_dice":  handleRollDice,
	"game_state": handleGameState,
}

func handleChatMessage(s *Server, conn *websocket.Conn, data interface{}) {
	// Process the chat message (e.g., broadcast it to all players in the room)
	log.Println("Chat message received: ", data)

	sender := conn.RemoteAddr().String()
	log.Println("sender: ", sender);
	message := Message{
		Type: "chat",
		Data: map[string]interface{} {
			"sender": sender,
			"data": data,
		},
	}

	// Convert the message data to JSON
	jsonData, err := json.Marshal(message)
	if err != nil {
		log.Println("Error marshalling message:", err)
		return
	}

	// Iterate over all clients and send the message to them
	s.clientsMutex.Lock()
	defer s.clientsMutex.Unlock()

	for client := range s.clients {
		go func(client *websocket.Conn, msg []byte) {
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("Error sending message to client: ", err)
			}
		}(client, jsonData)
	}
}

func handleRollDice(s *Server, conn *websocket.Conn, data interface{}) {
	log.Println("Dice roll received")

	sender := conn.RemoteAddr().String()
	dice1 := rand.Intn(6) + 1
	dice2 := rand.Intn(6) + 1


	log.Println("sender: ", sender);
	message := Message{
		Type: "roll_dice",
		Data: map[string]interface{} {
			"sender": sender,
			"dice1": dice1,
			"dice2": dice2,
		},
	}

	// Convert the message data to JSON
	jsonData, err := json.Marshal(message)
	if err != nil {
		log.Println("Error marshalling message:", err)
		return
	}

	// Iterate over all clients and send the message to them
	s.clientsMutex.Lock()
	defer s.clientsMutex.Unlock()

	for client := range s.clients {
		go func(client *websocket.Conn, msg []byte) {
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("Error sending message to client: ", err)
			}
		}(client, jsonData)
	}
}

func handleGameState(s *Server, conn *websocket.Conn, data interface{}) {
	// Process the game state message (e.g., update game state, validate move)
	log.Println("Game state update received: ", data)
	// Update the game state, check for win conditions, etc.
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true }, // Allow any origin
}

func (s *Server) readLoop(conn *websocket.Conn) {
	defer s.removeClient(conn)

	for {
		// Read any message sent over the websocket connection
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// Unmarshal the JSON message into a struct
		var msg Message
		if err := json.Unmarshal(p, &msg); err != nil {
			log.Println("JSON Unmarshal error:", err)
			continue
		}

		// Handle the message based on its type
		if handler, exists := messageHandlers[msg.Type]; exists {
			handler(s, conn, msg.Data)
		} else {
			log.Println("Unknown message type:", msg.Type)
		}
	}
}

// Functions to add and remove clients
func (s *Server) addClient(conn *websocket.Conn) {
	s.clientsMutex.Lock()
	defer s.clientsMutex.Unlock()

	s.clients[conn] = true
}

func (s *Server) removeClient(conn *websocket.Conn) {
	s.clientsMutex.Lock()
	defer s.clientsMutex.Unlock()
	
	delete(s.clients, conn)
}

// Home endpoint just prints
func (s *Server) homeEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home page")
}

// Websocket endpoint to upgrade HTTP connection to WebSocket
func (s *Server) wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// Upgrade the connection from HTTP to WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Client successfully connected:", ws.RemoteAddr())
	s.addClient(ws)

	// Start reading from the WebSocket
	s.readLoop(ws)
}

func setupRoutes(s *Server) {
	http.HandleFunc("/", s.homeEndpoint)
	http.HandleFunc("/ws", s.wsEndpoint)
}

func main() {
	fmt.Println("Go WebSockets Server Running...")
	server := NewServer()
	setupRoutes(server)
	log.Fatal(http.ListenAndServe(":8080", nil))
}