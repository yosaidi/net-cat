package main

import (
	"netcat/server"
	"os"
)

func main() {

	port := ""
	ip := ""
	if len(os.Args) ==1 {
		port = "8989"
	}
	if len(os.Args) == 2  {
		port = os.Args[1]
	} else if len(os.Args)==3 {
		ip = os.Args[2]
		port = os.Args[1]
	}
    

	server := server.NewServer_(ip, port)
	server.Start()

}
