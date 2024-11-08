package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/MiddeNg/go-network-programming/server"
	"math"
	"net"
)

func main() {
	server.StartTCPServer(primeTime)
}

type request struct {
	Method string `json:"method"`
	Number int    `json:"number"`
}

func primeTime(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024*1024)
	for {
		var req request
		n, err := conn.Read(buf)
		fmt.Println("read", n, "bytes")
		if err != nil {
			if err.Error() == "EOF" {
				return
			}
			//panic(err)
		}
		requests := bytes.Split(buf[:n], []byte{'\n'})
		for _, request := range requests {
			if len(request) == 0 {
				continue
			}
			println("request: "+string(request), len(request))

			if err := json.Unmarshal(request, &req); err != nil {
				fmt.Println("json error: " + err.Error())
				conn.Write([]byte(err.Error()))
				//return

			}
			if req.Method != "isPrime" {
				fmt.Println("Method not found")
				conn.Write([]byte("Method not found"))
				//return

			}

			inInt := req.Number

			fmt.Printf("float: %f, int: %d\n", req.Number, inInt)
			if isPrime(inInt) {
				//example : {"method":"isPrime","prime":false}
				res := "{\"method\":\"isPrime\",\"prime\":true}\n"
				fmt.Printf("Prime! num: %d, res: %s\n", inInt, res)
				if _, err := conn.Write([]byte(res)); err != nil {
					fmt.Println(err)
					//panic(err)
				}
			} else {
				res := "{\"method\":\"isPrime\",\"prime\":false}\n"
				fmt.Printf("Not Prime! num: %d, res: %s\n", inInt, res)
				if _, err := conn.Write([]byte(res)); err != nil {
					//panic(err)
					fmt.Println(err)
				}
			}

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
