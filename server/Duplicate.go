package server

import (
	"bufio"
	"net"
)

func (s *Server) DuplicateName(conn net.Conn) (bool, string) {
	reader := bufio.NewReader(conn)
	name, err := reader.ReadString('\n')
	HandleError(err)

	for _, pseudo := range s.clients {

		if name[:len(name)-1] == pseudo.Pseudo {
			conn.Write([]byte("Name already taken, enter a new name: "))
			return false, name
		} else if len(name) == 1 {
			conn.Write([]byte("Empty name, enter a proper one: "))
			return false, name
		}

	}
	return true, name

}
