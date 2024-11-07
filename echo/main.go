package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	startTCPEchoServer()
}

func startTCPEchoServer() {
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
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
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
