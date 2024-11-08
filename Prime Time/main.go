package main

import (
	"encoding/json"
	"io"
	"math"
	"net"

	"github.com/MiddeNg/go-network-programming/server"
)

func main() {
	server.StartTCPServer(primeTime)
}

type request struct {
	Method string  `json:"method"`
	Number float32 `json:"number"`
}

func primeTime(conn net.Conn) {
	defer conn.Close()
	for {
		var req request
		mes, err := io.ReadAll(conn)
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal(mes, &req); err != nil {
			conn.Write([]byte(err.Error()))
			return
		}
		if req.Method != "isPrime" {
			conn.Write([]byte("Method not found"))
			return
		}
		inInt := int(req.Number)

		if isPrime(inInt) {
			//example : {"method":"isPrime","prime":false}
			conn.Write([]byte("{\"method\":\"isPrime\",\"prime\":true}"))
		} else {
			conn.Write([]byte("{\"method\":\"isPrime\",\"prime\":false}"))
		}
	}
}
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	maxDivisor := int(math.Floor(math.Sqrt(float64(n))))
	for i := 3; i <= maxDivisor; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
