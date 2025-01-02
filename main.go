package main

import (
	"fmt"
	"netcat/server"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Wrong number of arguments, usage: go run . [Port number] [IP adress]")
		return
	}

	port := os.Args[1]
	ip := ""
	if len(os.Args) == 2 {
		ip = "localhost"
	} else {
		ip = os.Args[2]
	}

	server := server.NewServer_(ip,port)
    server.Start()

	
}
