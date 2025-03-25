package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

/*
Define routes
*/


// Simple homepage for our server
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home page")
}

// A websocket is an upgraded version of a basic HTTP connection
// HTTP closes after each interaction, but websockets stay open, so its better for real-time stuff
// So the goal is to upgrade from HTTP conn to websocket
func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// Take any incoming http request, and accept it regardless of origin
	upgrader.CheckOrigin = func(r *htpp.Request) bool { return true }

	// Upgrade the connection (upgraded from basic http conn to websocket)
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client successfuly connected...")
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

// Basic main function. Simply just
// 1. Setup the routes
// 2. Listen on port 8080
func main() {
	fmt.Println("Go websockets")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}