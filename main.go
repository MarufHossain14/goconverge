package main

import (
	"log"
	"net"
)

func main(){
	S := newServer()
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
	defer listener.Close()
	log.Print("started server on :8888")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %s", err.Error())
			continue
		}
		go s.newClient(conn)
	}
}