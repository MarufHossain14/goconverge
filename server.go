// the central page to manage all the information

package main

import "log"

type server struct {
	rooms map[string]*room
	commands chan command // channel for commands to be processed

}

//below is a function  to initial a new server.
func newServer() *server {
	return &server{
		rooms: make(map[string]*room),
		commands: make(chan command, 100), // buffered channel for commands
	}
}

func (s *server) newClient(conn net.Conn){
	log.Printf("New client connected: %s", conn.RemoteAddr().String())
	c := &client{
		conn: conn,
		nick: "anonymous", // default nickname
		commands: s.commands
	}
	
}