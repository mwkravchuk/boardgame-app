package handlers

import (
	"boardgame-app/server/network"
	"boardgame-app/server/shared"
)

func BuyHouse(s *network.Server, sender *shared.Client, msg shared.Message) {
	room, ok := IsInValidRoom(s, sender)
	if !ok {
		return
	}

	propertyIdx := int(msg.Data.(float64))

	game := room.GameState
  player := game.Players[sender.Id]
  prop := &game.Properties[propertyIndex]


}

func SellHouse(s *network.Server, sender *shared.Client, msg shared.Message) {
	room, ok := IsInValidRoom(s, sender)
	if !ok {
		return
	}

	propertyIdx := int(msg.Data.(float64))


}

func ToggleMortgage(s *network.Server, sender *shared.Client, msg shared.Message) {
	room, ok := IsInValidRoom(s, sender)
	if !ok {
		return
	}

	propertyIdx := int(msg.Data.(float64))
}