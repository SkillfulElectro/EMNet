package server

import (
	"EMNet/handlers"
	"net"
	"sync"
)

var (
	Peers_IP = make(map[int]string)
	Mutex    sync.Mutex
)

type Job struct {
	Conn net.Conn
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	handlers.HandleConnection(conn)
}
