package main

import (
	"bufio"
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
	bufReader := bufio.NewReader(conn)

	var data []byte
	var err error
	for {
		if data, err = io.ReadAll(bufReader); err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
	}

	fmt.Printf("Received: %s\n", data)
	if _, err := conn.Write(data); err != nil {
		panic(err)
	}
}
