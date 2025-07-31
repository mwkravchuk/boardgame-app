package handlers

import (
	"log"
	"fmt"
	"boardgame-app/server/network"
	"boardgame-app/server/shared"   
)

// handler function template
type HandlerFunc func(*network.Server, *shared.Client, shared.Message)

var Registry = map[string]HandlerFunc{
	"start_game":       StartGame,
	"join_room":        JoinRoom,
	"create_room":      CreateRoom,
	"new_turn":         NewTurn,
	"new_id":           NewId,
	"chat":             Chat,
	"console":          Console,
	"roll_dice":        Roll,
	"color_selected":   ColorSelected,
	"buy_property":     BuyProperty,
	"auction_property": AuctionProperty,
	"pay_rent":         PayRent,
	"bankrupt":         Bankrupt,
	"propose_trade":    ProposeTrade,
	"respond_to_trade": RespondToTrade,

	"debug_give_property": DebugGiveProperty,
}

// dispatcher --> sends message to corresponding function
func Dispatch(s *network.Server, c *shared.Client, msg shared.Message) {
	if h, ok := Registry[msg.Type]; ok {
		h(s, c, msg)
	} else {
		fmt.Println("unknown message:", msg.Type)
	}
}

func ReadLoop(s *network.Server, c *shared.Client) {
	defer s.RemoveClient(c)

	for {
		var msg shared.Message
		if err := c.Conn.ReadJSON(&msg); err != nil {
			log.Println("read:", err)
			return
		}
		Dispatch(s, c, msg)
	}
}
