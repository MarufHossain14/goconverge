// the central page to manage all the information

package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

type server struct {
	rooms    map[string]*room
	commands chan command // channel for commands to be processed

}

// below is a function  to initial a new server.
func newServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command, 100), // buffered channel for commands
	}
}

func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NICK:
			s.nick(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_ROOMS:
			s.listRooms(cmd.client, cmd.args)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client, cmd.args)
		}
	}
}

func (s *server) newClient(conn net.Conn) {
	log.Printf("New client connected: %s", conn.RemoteAddr().String())
	c := &client{
		conn:     conn,
		nick:     "anonymous", // default nickname
		commands: s.commands,
	}
	c.readInput() // start reading input from the client
}

func (s *server) nick(c *client, args []string) {
	if len(args) < 2 {
		c.err(fmt.Errorf("nickname required. Usage: /nick <nickname>"))
		return
	}
	c.nick = args[1]
	c.msg(fmt.Sprintf("You are now known as %s", c.nick))
}
func (s *server) join(c *client, args []string) {
	if len(args) < 2 {
		c.err(fmt.Errorf("room name required. Usage: /join <roomname>"))
		return
	}
	roomName := args[1]
	r, ok := s.rooms[roomName]
	if !ok {
		r = &room{
			name:    roomName,
			members: make(map[net.Addr]*client),
		}
		s.rooms[roomName] = r
	}
	r.members[c.conn.RemoteAddr()] = c

	s.quitCurrentRoom(c) // remove from current room if any

	c.room = r

	r.broadcast(c, fmt.Sprintf("%s has joined the room", c.nick))
	c.msg(fmt.Sprintf("You have joined the room: %s", r.name))

}
func (s *server) listRooms(c *client, args []string) {
	var roomList []string
	for name := range s.rooms {
		roomList = append(roomList, name)
	}
	c.msg(fmt.Sprintf("Available rooms: %s", strings.Join(roomList, ", ")))
}
func (s *server) msg(c *client, args []string) {
	if len(args) < 2 {
		c.err(fmt.Errorf("message required. Usage: /msg <message>"))
		return
	}
	if c.room == nil {
		c.err(errors.New("you must join the room first"))
		return
	}
	c.room.broadcast(c, c.nick+": "+strings.Join(args[1:], " "))

}
func (s *server) quit(c *client, args []string) {
	log.Printf("client has disconnected: %s", c.conn.RemoteAddr().String())
	s.quitCurrentRoom(c) // remove from current room if any

	c.msg("sad to see you go")
	c.conn.Close() // close the connection
}

func (s *server) quitCurrentRoom(c *client) {
	if c.room != nil {
		delete(c.room.members, c.conn.RemoteAddr())
		c.room.broadcast(c, fmt.Sprintf("%s has left the room", c.nick))
	}

}
