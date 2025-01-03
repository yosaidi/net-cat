package server

import (
	"bufio"
)

func (s *Server) ClientConnection(client Client) {
	defer client.conn.Close()

	reader := bufio.NewReader(client.conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			client = s.Broadcast(client, client.Pseudo, "leave")
			break
		}
		client = s.Broadcast(client, message, "message")

	}

}
