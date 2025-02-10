package main

import (
	"fmt"
	"log"
	"net-cat/server"
	"os"
)

func main() {
	port := "8989"

	if len(os.Args) > 2 {
		log.Println("[USAGE]: ./TCPChat $port")
		return
	}
	if len(os.Args) == 2 {
		port = os.Args[1]
	}

	s, err := server.NewServer(port,10)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer s.Ln.Close()
	fmt.Println("Starting the server at localhost:", port)

	if err := s.RunServer(); err != nil {
		fmt.Println(err)
		return
	}
}
