package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/websocket"
	"sync"
)

type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

var clients = make(map[*websocket.Conn]bool) // store active connections
var clientsMutex sync.Mutex                 // Mutex to protect access to `clients`

var messageHandlers = map[string]func(*websocket.Conn, Message){
	"chat":      handleChatMessage,
	"game_state": handleGameState,
}

func handleChatMessage(conn *websocket.Conn, msg Message) {
	// Process the chat message (e.g., broadcast it to all players in the room)
	log.Println("Chat message received: ", msg.Data)

	// Iterate over all clients and send the message to them
	clientsMutex.Lock() // Lock the mutex before accessing clients map
	defer clientsMutex.Unlock()

	for client := range clients {
		go func(client *websocket.Conn) {
			err := client.WriteMessage(websocket.TextMessage, msg.Data)
			if err != nil {
				log.Println("Error sending message to client: ", err)
			}
		}(client)
	}
}

func handleGameState(conn *websocket.Conn, msg Message) {
	// Process the game state message (e.g., update game state, validate move)
	log.Println("Game state update received: ", msg.Data)
	// Update the game state, check for win conditions, etc.
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true }, // Allow any origin
}

func reader(conn *websocket.Conn) {
	defer func() {
		clientsMutex.Lock()
		delete(clients, conn) // Remove the client from the map when it disconnects
		clientsMutex.Unlock()
		conn.Close()
	}()

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
			handler(conn, msg)
		} else {
			log.Println("Unknown message type:", msg.Type)
		}
	}
}

// Simple homepage for our server
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home page")
}

// A websocket endpoint handler to upgrade HTTP connection to WebSocket
func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// Upgrade the connection from HTTP to WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Client successfully connected:", ws.RemoteAddr())

	clientsMutex.Lock()
	clients[ws] = true // Add the new client to the map
	clientsMutex.Unlock()

	// Start reading from the WebSocket
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Go WebSockets Server Running...")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}