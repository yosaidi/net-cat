package server

import (
	"fmt"
	"net"
)

func NewServer_(ip, port string) *Server {
	return &Server{
		IP:    ip,
		PORT:  port,
		quich: make(chan struct{}),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen(Type, fmt.Sprintf("%s:%s", s.IP, s.PORT))
	HandleError(err)
	defer ln.Close()
	fmt.Println("Server started !")
	s.ln = ln

	go s.Accept()

	<-s.quich

	return nil
}

func (s *Server) Accept() {
	for {
		conn, err := s.ln.Accept()
		HandleError(err)
		fmt.Println("New connection from :", conn.RemoteAddr().String())
		client := Client{
			conn: conn,
		}

		if len(s.clients) == 10 {
			client.conn.Write([]byte("Server is full, 10 Users already connected.\n"))
			client = Client{}
		} else {
			ascii := AsciiArt()
			client.conn.Write([]byte(ascii + "Enter your name: "))
			go s.Read(conn)
		}
	}
}

func (s *Server) Read(conn net.Conn) {
	defer conn.Close()
	
	for {
		
		
	}
}
