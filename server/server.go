package server

import (
	"net"
	"net-cat/clients"
	"sync"
)

type Server struct {
	Ln   net.Listener
	Addr string
}

func NewServer(adress string) (*Server, error) {
	listener, err := net.Listen("tcp", "localhost:"+adress)

	if err != nil {
		return nil, err
	}

	return &Server{
		Ln:   listener,
		Addr: adress,
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
		conn, err := s.Ln.Accept()
		if err != nil {
			return err
		}
		newclient := clients.NewClient(conn)
		go newclient.HandleClient(statusch, messagech, chat)
	}
}
