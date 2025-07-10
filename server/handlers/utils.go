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

func InitializeProperties() []shared.Property {
	properties := make([]shared.Property, 40)
	properties[1] = shared.Property{
		Name: "Mediterranean Avenue",
		Color: "brown",
		Price: 60,
		Rent: 2,
		OwnerID: "",
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[3] = shared.Property{
		Name: "Baltic Avenue",
		Color: "brown",
		Price: 60,
		Rent: 4,
		OwnerID: "",
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[5] = shared.Property{
		Name: "Reading Railroad",
		Color: "black",
		Price: 200,
		Rent: 25,
		OwnerID: "",
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[6] = shared.Property{
		Name: "Oriental Avenue",
		Color: "light blue",
		Price: 100,
		Rent: 6,
		OwnerID: "",
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[8] = shared.Property{
		Name: "Vermont Avenue",
		Color: "light blue",
		Price: 100,
		Rent: 6,
		OwnerID: "",
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[9] = shared.Property{
		Name: "Connecticut Avenue",
		Color: "light blue",
		Price: 120,
		Rent: 8,
		OwnerID: "",
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[11] = shared.Property{
		Name: "St. Charles Place",
		Color: "pink",
		Price: 140,
		Rent: 10,
		OwnerID: "",
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[12] = shared.Property{
		Name: "Electric Company",
		Color: "white",
		Price: 150,
		Rent: 0,
		OwnerID: "",
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[13] = shared.Property{
		Name: "States Avenue",
		Color: "pink",
		Price: 140,
		Rent: 10,
		OwnerID: "",
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[14] = shared.Property{
		Name: "Virginia Avenue",
		Color: "pink",
		Price: 160,
		Rent: 12,
		OwnerID: "",
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[15] = shared.Property{
		Name: "Pennsylvania Railroad",
		Color: "black",
		Price: 200,
		Rent: 25,
		OwnerID: "",
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[16] = shared.Property{
		Name: "St. James Place",
		Color: "orange",
		Price: 180,
		Rent: 14,
		OwnerID: "",
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[18] = shared.Property{
		Name: "Tennessee Avenue",
		Color: "orange",
		Price: 180,
		Rent: 14,
		OwnerID: "",
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[19] = shared.Property{
		Name: "New York Avenue",
		Color: "orange",
		Price: 200,
		Rent: 16,
		OwnerID: "",
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[21] = shared.Property{
		Name: "Kentucky Avenue",
		Color: "red",
		Price: 220,
		Rent: 18,
		OwnerID: "",
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[23] = shared.Property{
		Name: "Indiana Avenue",
		Color: "red",
		Price: 220,
		Rent: 18,
		OwnerID: "",
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[24] = shared.Property{
		Name: "Illinois Avenue",
		Color: "red",
		Price: 240,
		Rent: 20,
		OwnerID: "",
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[25] = shared.Property{
		Name: "B & O Railroad",
		Color: "black",
		Price: 200,
		Rent: 25,
		OwnerID: "",
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[26] = shared.Property{
		Name: "Atlantic Avenue",
		Color: "yellow",
		Price: 260,
		Rent: 22,
		OwnerID: "",
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[27] = shared.Property{
		Name: "Ventnor Avenue",
		Color: "yellow",
		Price: 260,
		Rent: 22,
		OwnerID: "",
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[28] = shared.Property{
		Name: "Water Works",
		Color: "white",
		Price: 150,
		Rent: 0,
		OwnerID: "",
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[29] = shared.Property{
		Name: "Marvin Gardens",
		Color: "yellow",
		Price: 280,
		Rent: 24,
		OwnerID: "",
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[31] = shared.Property{
		Name: "Pacific Avenue",
		Color: "green",
		Price: 300,
		Rent: 26,
		OwnerID: "",
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[32] = shared.Property{
		Name: "North Carolina Avenue",
		Color: "green",
		Price: 300,
		Rent: 26,
		OwnerID: "",
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[34] = shared.Property{
		Name: "Pennsylvania Avenue",
		Color: "green",
		Price: 320,
		Rent: 28,
		OwnerID: "",
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[35] = shared.Property{
		Name: "Short Line",
		Color: "black",
		Price: 200,
		Rent: 25,
		OwnerID: "",
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[37] = shared.Property{
		Name: "Park Place",
		Color: "blue",
		Price: 350,
		Rent: 35,
		OwnerID: "",
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
	}
	properties[39] = shared.Property{
		Name: "Boardwalk",
		Color: "blue",
		Price: 400,
		Rent: 50,
		OwnerID: "",
		IsProperty: true,
		IsOwned: false,
		IsMortgaged: false,
	}

	return properties
}