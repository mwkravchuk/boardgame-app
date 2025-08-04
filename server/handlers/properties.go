package handlers

import (
	"fmt"
	"boardgame-app/server/network"
	"boardgame-app/server/shared"
)

func BuyHouse(s *network.Server, sender *shared.Client, msg shared.Message) {
	room, ok := IsInValidRoom(s, sender)
	if !ok {
		return
	}

	s.BroadcastToRoom(room, shared.Message{
		Type: "game_state",
		Data: room.GameState,
	})

}

func SellHouse(s *network.Server, sender *shared.Client, msg shared.Message) {
	room, ok := IsInValidRoom(s, sender)
	if !ok {
		return
	}

	//propertyIdx := int(msg.Data.(float64))

	s.BroadcastToRoom(room, shared.Message{
		Type: "game_state",
		Data: room.GameState,
	})

}

func ToggleMortgage(s *network.Server, sender *shared.Client, msg shared.Message) {

	fmt.Println("toggling mortgage")

	room, ok := IsInValidRoom(s, sender)
	if !ok {
		return
	}

	player := room.GameState.Players[sender.Id]
	propertyIdx := int(msg.Data.(float64))
	fmt.Println("propertyIdx: ", propertyIdx)
	property := &room.GameState.Properties[propertyIdx]

	if !ownsProperty(player, propertyIdx) || property.NumHouses > 0 {
		fmt.Println("Trying to toggle mortgage: Does not own property or there are houses on the prop")
		return
	}

	if property.IsMortgaged {
		cost := int(float64(property.Price) * 0.55)
		if player.Money < cost {
			fmt.Println("not enough money to unmortgage")
			return
		}
		player.Money -= cost
		property.IsMortgaged = false
	} else {
		player.Money += property.Price / 2
		property.IsMortgaged = true
	}

	s.BroadcastToRoom(room, shared.Message{
		Type: "game_state",
		Data: room.GameState,
	})
}

func ownsProperty(player *shared.PlayerState, propertyIdx int) bool {
	for _, idx := range player.PropertiesOwned {
		if idx == propertyIdx {
			return true
		}
	}
	return false
}