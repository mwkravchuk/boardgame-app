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

	propertyIdx := int(msg.Data.(float64))
	player := room.GameState.Players[sender.Id]

	okToBuild, reason := CanBuildHouse(room.GameState, propertyIdx, player)
	if !okToBuild {
		fmt.Println("Cannot build house: ", reason)
		return
	}

	// Apply purchase
	property := &room.GameState.Properties[propertyIdx]
  player.Money -= property.HouseCost
  property.NumHouses++

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

	propertyIdx := int(msg.Data.(float64))
	player := room.GameState.Players[sender.Id]

	okToSellHouse, reason := CanSellHouse(room.GameState, propertyIdx, player)
	if !okToSellHouse {
		fmt.Println("Cannot sell house: ", reason)
		return
	}

	// Apply purchase
	property := &room.GameState.Properties[propertyIdx]
  player.Money += property.HouseCost / 2
	property.NumHouses--

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

// CalculateRent returns the current rent amount for a property given the player's state and the game state.
// Handles utilities, railroads, and colored properties with/without monopoly and houses.
func CalculateRent(player *shared.PlayerState, propertyIdx int, gameState *shared.GameState) int {
	property := &gameState.Properties[propertyIdx]

	if property.IsMortgaged {
		return 0
	}

	switch property.Color {
	// ---- RAILROADS ----
	case "black":
		count := 0
		for _, idx := range player.PropertiesOwned {
			if gameState.Properties[idx].Color == "black" {
				count++
			}
		}
		if count == 0 {
			return property.DefaultRent
		}
		if count > len(property.RentStages) {
			count = len(property.RentStages)
		}
		return property.RentStages[count-1]

	// ---- UTILITIES ----
	case "white":
		// Utility rent is roll value * multiplier
		rollValue := gameState.LastRoll

		count := 0
		for _, idx := range player.PropertiesOwned {
			if gameState.Properties[idx].Color == "black" {
				count++
			}
		}

		multiplier := 4
		if count == 2 {
			multiplier = 10
		}

		return rollValue * multiplier

	// ---- REGULAR PROPERTIES ----
	default:
		if ownsMonopoly(gameState, player, property.Color) {
			return property.RentStages[property.NumHouses] // Takes care of 0-5 houses.
		}
		// No monopoly: use default rent
		return property.DefaultRent
	}
}

func CanBuildHouse(game *shared.GameState, propertyIdx int, player *shared.PlayerState) (bool, string) {
	property := game.Properties[propertyIdx]

	// Must own property
	if !ownsProperty(player, propertyIdx) {
		return false, "Player does not own property"
	}

	// Property cannot be mortgaged
	if property.IsMortgaged {
		return false, "Cannot build on mortgaged property"
	}

	// Must own monopoly
	if !ownsMonopoly(game, player, property.Color) {
		return false, "Cannot build; you don't own the full set"
	}

	// Cannot exceed 5 houses
	if property.NumHouses >= 5 {
		return false, "Cannot build more than 5 houses (hotel)"
	}

	// Must have enough money
	if player.Money < property.HouseCost {
		return false, "Not enough money to buy house"
	}

	// Even building rule
	group := game.ColorGroups[property.Color]
	minHouses := game.Properties[group[0]].NumHouses
	for _, idx := range group {
		if game.Properties[idx].NumHouses < minHouses {
			minHouses = game.Properties[idx].NumHouses
		}
	}

	if property.NumHouses > minHouses {
		return false, "Must build evenly across properties"
	}

	return true, ""
}

func CanSellHouse(game *shared.GameState, propertyIdx int, player *shared.PlayerState) (bool, string) {
	property := game.Properties[propertyIdx]

	if !ownsProperty(player, propertyIdx) {
		return false, "Player does not own property"
	}

	if property.IsMortgaged {
		return false, "Cannot sell house on mortgaged property"
	}

	if !ownsMonopoly(game, player, property.Color) {
		return false, "Cannot build; you don't own the full set"
	}

	if property.NumHouses == 0 {
		return false, "No houses to sell"
	}

	// Even selling rule
	group := game.ColorGroups[property.Color]
	maxHouses := game.Properties[group[0]].NumHouses
	for _, idx := range group {
		if game.Properties[idx].NumHouses > maxHouses {
			maxHouses = game.Properties[idx].NumHouses
		}
	}

	if property.NumHouses < maxHouses {
		return false, "Must sell houses evenly across properties"
	}

	return true, ""
}

