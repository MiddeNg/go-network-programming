package server

import (
	"fmt"
	"net"
)

func StartTCPServer(handler func(conn net.Conn)) {
	listener, err := net.Listen("tcp", ":10000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Listening on :10000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("Accepted connection")
		go handler(conn)
	}
}
