package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// bootstrap
func main() {
	hub := NewServer()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("upgrade:", err)
			return
		}
		hub.AddClient(conn)
	})

	log.Println("Websocket Server Running")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// gorilla upgrader lives here so server.go can reuse it
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}
