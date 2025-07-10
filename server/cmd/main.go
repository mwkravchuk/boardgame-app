package main

import (
	"log"
	"net/http"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"boardgame-app/server/network"
	"boardgame-app/server/shared"
	"boardgame-app/server/handlers"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var server = network.NewServer()

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Websocket upgrade failed:", err)
		return
	}

	client := &shared.Client{
		Conn: conn,
		Id: uuid.NewString(),
	}

	server.AddClient(client)
	handlers.ReadLoop(server, client)
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
