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
	Clients   map[string]*Client
	GameState *GameState
}

type GameState struct {
	Players      map[string]*PlayerState `json:"players"`
	TurnOrder    []string 							 `json:"turnOrder"`
	CurrentTurn  int 										 `json:"currentTurn"`
	BoardState   []int 									 `json:"boardState"`
	Properties   []Property							 `json:"properties"`
	CurrentTrade *TradeOffer             `json:"tradeOffer"`
	ColorGroups  map[string][]int        `json:"colorGroups"` // "brown" -> {1, 3}
	LastRoll     int                     `json:"lastRoll"`
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

type Property struct {
	Name 				string `json:"name"`
	Color       string `json:"color"`
	Price 	    int    `json:"price"`
	DefaultRent int		 `json:"defaultRent"`
	IsProperty  bool   `json:"isProperty"`
	IsOwned     bool   `json:"isOwned"`
	NumHouses   int    `json:"numHouses"`
  RentStages  []int  `json:"rentStages"` // rent for 0, 1, ..., 5 houses
	HouseCost   int    `json:"houseCost"`
	IsMortgaged bool   `json:"isMortgaged"`
	OwnerID     string `json:"ownerId"`
}

type ProposeTradePayload struct {
	TargetID        string   `json:"targetId"`
  MyOfferMoney    int      `json:"myOfferMoney"`
	MyOfferProps    []int    `json:"myOfferProps"`
  TheirOfferMoney int      `json:"theirOfferMoney"`
  TheirOfferProps []int    `json:"theirOfferProps"`
}

type TradeOffer struct {
	FromPlayerID string `json:"fromPlayerId`
  ToPlayerID   string `json:"toPlayerId"`
  OfferMoney   int    `json:"offerMoney,omitempty"`
  OfferProps   []int  `json:"offerProps"`
  RequestMoney int    `json:"requestMoney,omitempty"`
  RequestProps []int	`json:"requestProps"`
  Status       string `json:"status"`// "pending", "accepted", "rejected"
}