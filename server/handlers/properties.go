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
	if !ownsProperty(player, propertyIdx) {
    fmt.Println("Player does not own property")
    return
	}

	property := &room.GameState.Properties[propertyIdx]
	if property.IsMortgaged {
    fmt.Println("Cannot build on mortgaged property")
    return
	}

	// Must own full color set
	color := property.Color
	ownsAll := true
	for i, p := range room.GameState.Properties {
		if p.Color == color && !ownsProperty(player, i) {
			ownsAll = false
			break
		}
	}
	if !ownsAll {
		fmt.Println("Player does not own full monopoly")
		return
	}

	// Balanced building check
	var groupProps []*shared.Property
	for i := range room.GameState.Properties {
		if room.GameState.Properties[i].Color == color {
			groupProps = append(groupProps, &room.GameState.Properties[i])
		}
	}

	minHouses := groupProps[0].NumHouses
	for _, gp := range groupProps {
		if gp.NumHouses < minHouses {
			minHouses = gp.NumHouses
		}
	}

	if property.NumHouses > minHouses {
		fmt.Println("Cannot build here; must build evenly")
		return
	}

	// Funds check
	if player.Money < property.HouseCost {
		fmt.Println("Not enough money")
		return
	}

	// Apply purchase
  player.Money -= property.HouseCost
  property.NumHouses++

	// TO DO CALCULATE RENT SOMEHOW>
  //property.CurrentRent = property.RentStages[property.NumHouses] // assume RentStages exists

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

func ownsProperty(player *shared.PlayerState, propertyIdx int) bool {
	for _, idx := range player.PropertiesOwned {
		if idx == propertyIdx {
			return true
		}
	}
	return false
}

func ownsMonopoly(game *shared.GameState, player *shared.PlayerState, color string) bool {
	// Get the property indices for this color group
	group, exists := game.ColorGroups[color]
	if !exists {
		return false
	}

	// Check that player owns all properties in the group
	for _, propertyIdx := range group {
		if !ownsProperty(player, propertyIdx) {
			return false
		}
	}

	return true
}

