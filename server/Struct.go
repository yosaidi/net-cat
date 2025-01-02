package server

import (
	"net"
	"sync"
)

type Server struct {
	ln net.Listener
	clients []Client
	mutex   sync.Mutex
	IP      string
	PORT    string
	quich   chan struct{}
}


type Client struct{
    Pseudo string
     conn net.Conn

}

const (
	Type = "tcp"
)