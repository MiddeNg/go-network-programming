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
	port := 10000
	listener, err := net.Listen("tcp", ":"+fmt.Sprint(port))
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Listening on :" + fmt.Sprint(port))

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
		if err.Error() == "EOF" {
			fmt.Println("reached EOF")
		}
		panic(err)
	}
	fmt.Println("Read", len(data), "bytes")
	fmt.Printf("Received: %s\n", data)

	if _, err := conn.Write(data); err != nil {
		panic(err)
	}
}
