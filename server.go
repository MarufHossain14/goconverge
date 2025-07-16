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

func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		CMD_NICK:
			s.nick(cmd.client, cmd.args)
		CMD_JOIN:
			s.join(cmd.client, cmd.args)
		CMD_ROOMS:
			s.listRooms(cmd.client, cmd.args)
		CMD_MSG:
			s.msg(cmd.client, cmd.args)
		CMD_QUIT:
			s.quit(cmd.client, cmd.args)
			
		}
	}
}

func (s *server) newClient(conn net.Conn){
	log.Printf("New client connected: %s", conn.RemoteAddr().String())
	c := &client{
		conn: conn,
		nick: "anonymous", // default nickname
		commands: s.commands
	}
	c.readInput() // start reading input from the client
}