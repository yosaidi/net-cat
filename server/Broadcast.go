package server

import (
	"fmt"
	"time"
)

func (s *Server) Broadcast(sender Client, message, status string) Client {

	if status == "joined" {
		for _, client := range s.clients {
			client.conn.Write([]byte("\033[32m" + time.Now().Format("2006-01-02 15:04:05") + "] " + message + " has joined the chat.\n" + "\033[0m"))

		}

	} else if status == "leaved" {
		for _, client := range s.clients {
			client.conn.Write([]byte("\033[32m" + time.Now().Format("2006-01-02 15:04:05") + "] " + message + " has leaved the chat.\n" + "\033[0m"))

		}
	} else if status == "message" {
		for _, client := range s.clients {
			client.conn.Write([]byte("\033[37m" + "[" + time.Now().Format("2006-01-02 15:04:05") + "]" + "[" + "\033[36m" + string(sender.Pseudo) + "\033[0m" + "]:" + message))
			fmt.Println(client.Pseudo)
		}

	}
	return sender
}
