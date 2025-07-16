package main

import "net"

type room struct {
	name    string // name of the room
	members map[net.Addr]*client // map of clients in the room, key is the client's connection address

}