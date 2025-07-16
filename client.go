//client will be focuse on keeping information about the user, such as name and current tcp connection, and also current room.

package main

import (
	"net"
	"strings"
)

type client struct{
	conn net.Conn // connection to the server
	name string // name of the user, if not set, user will be anonymous
	room *room // current room user is in, if not set, user is not in any room
	commands chan command // channel for commands to be processed
}

func (c *client) readInput() {
	// This function will read input from the client and send commands to the server
	// Implementation will depend on how you want to handle input (e.g., using bufio.Scanner)
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\r\n")

		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		switch cmd{
		case "/nick":
		case "/join":
		case "/rooms":
		case "/msg":
		case "/quit":
		default:
			c.err(fmt.Errorf("unknown command: %s", cmd))
				
		}
	}


}

func (c *client) err(err error) {
	c.conn.Write([]byte("Err: " + err.Error() + "\n"))
}
func (c *client) msg(msg string) {
	c.conn.Write([]byte(">: " + msg + "\n"))
}
