package network

import (
	"log"
	"boardgame-app/server/shared"
)

func (s *Server) AddClient(c *shared.Client) {
	s.ClientsMutex.Lock()
	defer s.ClientsMutex.Unlock()
	s.ClientsById[c.Id] = c
	log.Println("Added client: ", c.Id)
}

func (s *Server) RemoveClient(c *shared.Client) {
	s.ClientsMutex.Lock()
	defer s.ClientsMutex.Unlock()
	delete(s.ClientsById, c.Id)
	log.Println("Removed client: ", c.Id)
}