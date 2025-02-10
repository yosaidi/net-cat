package server

import (
	"fmt"
	"net"
	"net-cat/clients"
	"sync"
)

type Server struct {
	Ln         net.Listener
	Addr       string
	MaxClients int
	mu         sync.Mutex
	connCount  int
}

func NewServer(address string, maxClients int) (*Server, error) {
	listener, err := net.Listen("tcp", ":"+address)
	if err != nil {
		return nil, err
	}

	return &Server{
		Ln:         listener,
		Addr:       address,
		MaxClients: maxClients,
	}, nil
}

func (s *Server) RunServer() error {
	statusch := make(chan clients.ConnectionStatus)
	messagech := make(chan clients.BroadcastMessage)
	var (
		muClients, muChat sync.Mutex
	)
	allclients := clients.NewClients(&muClients)
	chat := clients.NewChat(&muChat, allclients)
	go chat.HandleChatRoutine(statusch, messagech)

	for {
		s.mu.Lock()
		if s.connCount >= s.MaxClients {

			Conn, err := s.Ln.Accept()
			clients.FullGroup(Conn)
			if err != nil {
				fmt.Println(err)
			}
			Conn.Close()
		}
		s.mu.Unlock()

		conn, err := s.Ln.Accept()
		if err != nil {
			return err
		}

		s.mu.Lock()
		s.connCount++
		s.mu.Unlock()

		newclient := clients.NewClient(conn)
		go func() {
			newclient.HandleClient(statusch, messagech, chat)
			s.mu.Lock()
			s.connCount--
			s.mu.Unlock()
		}()
	}
}
