package main

import (
	"fmt"
	"net/http"
	"log"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "websocket endpoint")
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Go websockets")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}