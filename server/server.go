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
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("New connection from :", conn.RemoteAddr().String())
		client := Client{
			conn: conn,
		}

		if len(s.clients) == 10 {
			client.conn.Write([]byte("Server is full, 10 Users already connected.\n"))
			client = Client{}
		} else {
			ascii := AsciiArt()
			client.conn.Write([]byte(ascii))

			connectedusers = ""

			for _, connected := range s.clients {
				connectedusers += connected.Pseudo + ","
			}
			client.conn.Write([]byte("Welcome\n"))

			if len(connectedusers) == 0 {
				client.conn.Write([]byte("Server empty\n"))
			} else {
				client.conn.Write([]byte("Clients connected: " + "\033[34m" + connectedusers[:len(connectedusers)-2] + "\033[0m" + "\n"))
			}
			client.conn.Write([]byte("Enter your name: "))

			duplicate, name := s.DuplicateName(conn)

			for !duplicate {
				duplicate, name = s.DuplicateName(conn)
			}
			client = s.Broadcast(client, name[:len(name)-1], "joined")

			client = Client{
				conn:   conn,
				Pseudo: name[:len(name)-1],
			}
			s.mutex.Lock()
			s.clients = append(s.clients, client)
			s.mutex.Unlock()
			fmt.Println("Number of clients connected: ", len(s.clients))
			go s.ClientConnection(client)
		}
	}
}
