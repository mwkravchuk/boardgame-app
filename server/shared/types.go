package shared

import (
	"github.com/gorilla/websocket"
	"sync"
)

type Message struct {
	Type   string      `json:"type"`
	Sender string      `json:"sender,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

type Client struct {
	Conn *websocket.Conn
	Id   string
	Mu   sync.Mutex
}

type GameRoom struct {
	Code      string
	PlayerIDs map[string]bool
	GameState *GameState
}

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
	HasRolled       bool       `json:"hasRolled"`
	Color           string 		 `json:"color"`
	PropertiesOwned []int      `json:"properties"` // store indices of owned properties
}