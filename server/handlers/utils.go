package handlers

import (
	"fmt"
	"time"
	"math/rand"
	"boardgame-app/server/network"
	"boardgame-app/server/shared"
)

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const codeLength = 4

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateRoomCode() string {
	b := make([]byte, codeLength)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func IsInValidRoom(s *network.Server, sender *shared.Client) (*shared.GameRoom, bool) {
	roomCode, ok := s.ClientToRoomCode[sender.Id]
	if !ok {
		fmt.Println("Client not in a room")
		return nil, false
	}

	room, ok := s.Rooms[roomCode]
	if !ok {
		fmt.Println("Room not found: ", roomCode)
		return nil, false
	}

	return room, true
}

// Helper for property transfer
func transferProperties(props []int, from *shared.PlayerState, to *shared.PlayerState) {
	for _, propID := range props {
		// Remove from `from`
		newOwned := []int{}
		for _, p := range from.PropertiesOwned {
			if p != propID {
				newOwned = append(newOwned, p)
			}
		}
		from.PropertiesOwned = newOwned

		// Add to `to`
		to.PropertiesOwned = append(to.PropertiesOwned, propID)
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

func ownsMonopoly(gameState *shared.GameState, player *shared.PlayerState, color string) bool {
	// Get the property indices for this color group
	group, exists := gameState.ColorGroups[color]
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

func InitializeProperties() []shared.Property {
	properties := make([]shared.Property, 40)
	properties[1] = shared.Property{
		Name: "Mediterranean Avenue",
		Color: "brown",
		Price: 60,
		DefaultRent: 2,
		RentStages: []int{4, 10, 30, 90, 160, 250},
		NumHouses: 0,
		HouseCost: 50,
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
		OwnerID: "",
	}
	properties[3] = shared.Property{
		Name: "Baltic Avenue",
		Color: "brown",
		Price: 60,
		DefaultRent: 4,
		RentStages: []int{8, 20, 60, 180, 320, 450},
		NumHouses: 0,
		HouseCost: 50,
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
		OwnerID: "",
	}
	properties[5] = shared.Property{
		Name: "Reading Railroad",
		Color: "black",
		Price: 200,
		DefaultRent: 25,
		// TO DO
		RentStages: nil,
		NumHouses: 0,
		HouseCost: 0,
		OwnerID: "",
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
	}

	properties[6] = shared.Property{
		Name: "Oriental Avenue",
		Color: "light blue",
		Price: 100,
		DefaultRent: 6,
		RentStages: []int{12, 30, 90, 270, 400, 550},
		NumHouses: 0,
		HouseCost: 50,
		OwnerID: "",
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
	}

	properties[8] = shared.Property{
		Name: "Vermont Avenue",
		Color: "light blue",
		Price: 100,
		DefaultRent: 6,
		RentStages: []int{12, 30, 90, 270, 400, 550},
		NumHouses: 0,
		HouseCost: 50,
		OwnerID: "",
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
	}

	properties[9] = shared.Property{
    Name:        "Connecticut Avenue",
    Color:       "light blue",
    Price:       120,
    DefaultRent: 8,
    RentStages:  []int{16, 40, 100, 300, 450, 600}, // doubled base rent, 1 house, 2 houses, ...
    NumHouses:   0,
    HouseCost:   50,
    OwnerID:     "",
    IsProperty:  true,
    IsOwned:     false,
    IsMortgaged: false,
	}

	properties[11] = shared.Property{
		Name:        "St. Charles Place",
		Color:       "pink",
		Price:       140,
		DefaultRent: 10,
		RentStages:  []int{20, 50, 150, 450, 625, 750},
		NumHouses:   0,
		HouseCost:   50,
		OwnerID:     "",
		IsProperty:  true,
		IsOwned:     false,
		IsMortgaged: false,
	}

	properties[12] = shared.Property{
		Name:        "Electric Company",
		Color:       "white",
		Price:       150,
		DefaultRent: 0,
		RentStages:  nil, // Utilities handled differently
		NumHouses:   0,
		HouseCost:   0,
		OwnerID:     "",
		IsProperty:  true,
		IsOwned:     false,
		IsMortgaged: false,
	}

	properties[13] = shared.Property{
		Name:        "States Avenue",
		Color:       "pink",
		Price:       140,
		DefaultRent: 10,
		RentStages:  []int{20, 50, 150, 450, 625, 750},
		NumHouses:   0,
		HouseCost:   50,
		OwnerID:     "",
		IsProperty:  true,
		IsOwned:     false,
		IsMortgaged: false,
	}

	properties[14] = shared.Property{
		Name:        "Virginia Avenue",
		Color:       "pink",
		Price:       160,
		DefaultRent: 12,
		RentStages:  []int{24, 60, 180, 500, 700, 900},
		NumHouses:   0,
		HouseCost:   50,
		OwnerID:     "",
		IsProperty:  true,
		IsOwned:     false,
		IsMortgaged: false,
	}

	properties[15] = shared.Property{
		Name:        "Pennsylvania Railroad",
		Color:       "black",
		Price:       200,
		DefaultRent: 25, // base rent for 1 railroad
		RentStages:  []int{25, 50, 100, 200},
		NumHouses:   0,
		HouseCost:   0,
		OwnerID:     "",
		IsProperty:  true,
		IsOwned:     false,
		IsMortgaged: false,
	}

	properties[16] = shared.Property{
		Name:        "St. James Place",
		Color:       "orange",
		Price:       180,
		DefaultRent: 14,
		RentStages:  []int{28, 70, 200, 550, 750, 950},
		NumHouses:   0,
		HouseCost:   100,
		OwnerID:     "",
		IsProperty:  true,
		IsOwned:     false,
		IsMortgaged: false,
	}

	properties[18] = shared.Property{
		Name:        "Tennessee Avenue",
		Color:       "orange",
		Price:       180,
		DefaultRent: 14,
		RentStages:  []int{28, 70, 200, 550, 750, 950},
		NumHouses:   0,
		HouseCost:   100,
		OwnerID:     "",
		IsProperty:  true,
		IsOwned:     false,
		IsMortgaged: false,
	}

	properties[19] = shared.Property{
		Name:        "New York Avenue",
		Color:       "orange",
		Price:       200,
		DefaultRent: 16,
		RentStages:  []int{32, 80, 220, 600, 800, 1000},
		NumHouses:   0,
		HouseCost:   100,
		OwnerID:     "",
		IsProperty:  true,
		IsOwned:     false,
		IsMortgaged: false,
	}

	properties[21] = shared.Property{
		Name:        "Kentucky Avenue",
		Color:       "red",
		Price:       220,
		DefaultRent: 18,
		RentStages:  []int{36, 90, 250, 700, 875, 1050},
		NumHouses:   0,
		HouseCost:   150,
		OwnerID:     "",
		IsProperty:  true,
		IsOwned:     false,
		IsMortgaged: false,
	}

	properties[23] = shared.Property{
		Name:        "Indiana Avenue",
		Color:       "red",
		Price:       220,
		DefaultRent: 18,
		RentStages:  []int{36, 90, 250, 700, 875, 1050},
		NumHouses:   0,
		HouseCost:   150,
		OwnerID:     "",
		IsProperty:  true,
		IsOwned:     false,
		IsMortgaged: false,
	}

	properties[24] = shared.Property{
		Name:        "Illinois Avenue",
		Color:       "red",
		Price:       240,
		DefaultRent: 20,
		RentStages:  []int{40, 100, 300, 750, 925, 1100},
		NumHouses:   0,
		HouseCost:   150,
		OwnerID:     "",
		IsProperty:  true,
		IsOwned:     false,
		IsMortgaged: false,
	}

	properties[25] = shared.Property{
		Name:        "B & O Railroad",
		Color:       "black",
		Price:       200,
		DefaultRent: 25,
		RentStages:  []int{25, 50, 100, 200},
		NumHouses:   0,
		HouseCost:   0,
		OwnerID:     "",
		IsProperty:  true,
		IsOwned:     false,
		IsMortgaged: false,
	}

	properties[26] = shared.Property{
		Name:        "Atlantic Avenue",
		Color:       "yellow",
		Price:       260,
		DefaultRent: 22,
		RentStages:  []int{44, 110, 330, 800, 975, 1150},
		NumHouses:   0,
		HouseCost:   150,
		OwnerID:     "",
		IsProperty:  true,
		IsOwned:     false,
		IsMortgaged: false,
	}

	properties[27] = shared.Property{
		Name:        "Ventnor Avenue",
		Color:       "yellow",
		Price:       260,
		DefaultRent: 22,
		RentStages:  []int{44, 110, 330, 800, 975, 1150},
		NumHouses:   0,
		HouseCost:   150,
		OwnerID:     "",
		IsProperty:  true,
		IsOwned:     false,
		IsMortgaged: false,
	}

	properties[28] = shared.Property{
		Name:        "Water Works",
		Color:       "white",
		Price:       150,
		DefaultRent: 0,
		RentStages:  nil, // Utilities handled separately
		NumHouses:   0,
		HouseCost:   0,
		OwnerID:     "",
		IsProperty:  true,
		IsOwned:     false,
		IsMortgaged: false,
	}

	properties[29] = shared.Property{
		Name:        "Marvin Gardens",
		Color:       "yellow",
		Price:       280,
		DefaultRent: 24,
		RentStages:  []int{48, 120, 360, 850, 1025, 1200},
		NumHouses:   0,
		HouseCost:   150,
		OwnerID:     "",
		IsProperty:  true,
		IsOwned:     false,
		IsMortgaged: false,
	}

	properties[31] = shared.Property{
		Name:        "Pacific Avenue",
		Color:       "green",
		Price:       300,
		DefaultRent: 26,
		RentStages:  []int{52, 130, 390, 900, 1100, 1275},
		NumHouses:   0,
		HouseCost:   200,
		OwnerID:     "",
		IsProperty:  true,
		IsOwned:     false,
		IsMortgaged: false,
	}

	properties[32] = shared.Property{
		Name:        "North Carolina Avenue",
		Color:       "green",
		Price:       300,
		DefaultRent: 26,
		RentStages:  []int{52, 130, 390, 900, 1100, 1275},
		NumHouses:   0,
		HouseCost:   200,
		OwnerID:     "",
		IsProperty:  true,
		IsOwned:     false,
		IsMortgaged: false,
	}

	properties[34] = shared.Property{
		Name:        "Pennsylvania Avenue",
		Color:       "green",
		Price:       320,
		DefaultRent: 28,
		RentStages:  []int{56, 150, 450, 1000, 1200, 1400},
		NumHouses:   0,
		HouseCost:   200,
		OwnerID:     "",
		IsProperty:  true,
		IsOwned:     false,
		IsMortgaged: false,
	}

	properties[35] = shared.Property{
		Name:        "Short Line",
		Color:       "black",
		Price:       200,
		DefaultRent: 25,
		RentStages:  []int{25, 50, 100, 200},
		NumHouses:   0,
		HouseCost:   0,
		OwnerID:     "",
		IsProperty:  true,
		IsOwned:     false,
		IsMortgaged: false,
	}

	properties[37] = shared.Property{
		Name:        "Park Place",
		Color:       "blue",
		Price:       350,
		DefaultRent: 35,
		RentStages:  []int{70, 175, 500, 1100, 1300, 1500},
		NumHouses:   0,
		HouseCost:   200,
		OwnerID:     "",
		IsProperty:  true,
		IsOwned:     false,
		IsMortgaged: false,
	}

properties[39] = shared.Property{
	Name:        "Boardwalk",
	Color:       "blue",
	Price:       400,
	DefaultRent: 50,
	RentStages:  []int{100, 300, 750, 925, 1100, 1275},
	NumHouses:   0,
	HouseCost:   200,
	OwnerID:     "",
	IsProperty:  true,
	IsOwned:     false,
	IsMortgaged: false,
}

	return properties
}

func InitializeColorGroups() map[string][]int {
	return map[string][]int{
		"brown":       {1, 3},
		"light blue":  {6, 8, 9},
		"pink":        {11, 13, 14},
		"orange":      {16, 18, 19},
		"red":         {21, 23, 24},
		"yellow":      {26, 27, 29},
		"green":       {31, 32, 34},
		"blue":        {37, 39},
	}
}