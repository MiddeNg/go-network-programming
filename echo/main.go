package main

import (
	"io"
	"net"

	"github.com/MiddeNg/go-network-programming/server"
)

func main() {
	server.StartTCPServer(echo)
}

func echo(conn net.Conn) {
	defer conn.Close()

	var data []byte
	var err error

	if data, err = io.ReadAll(conn); err != nil {
		panic(err)
	}

	if _, err := conn.Write(data); err != nil {
		panic(err)
	}
}
