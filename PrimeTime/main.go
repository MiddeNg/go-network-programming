package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/MiddeNg/go-network-programming/server"
	"io"
	"math"
	"net"
	"strings"
)

func main() {
	server.StartTCPServer(primeTime)
}

type request struct {
	Method string `json:"method"`
	Number int    `json:"number"`
}

type request2 struct {
	Method string  `json:"method"`
	Number float64 `json:"number"`
}

func primeTime(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 0, 1024*1024)
	for {
		var req request
		var dummy request2
		for {
			n, err := conn.Read(buf[len(buf):cap(buf)])
			buf = buf[:len(buf)+n]
			if err != nil {
				if err == io.EOF {
					return
				}
				panic(err)
			}
			if buf[len(buf)-1] == '\n' {
				break
			}
			if len(buf) == cap(buf) {
				buf = append(buf, 0)[:n]
			}
		}
		requests := bytes.Split(buf, []byte{'\n'})
		buf = buf[:0]
		for _, request := range requests {
			if len(request) == 0 {
				continue
			}
			//println("request: "+string(request), len(request))

			if err := json.Unmarshal(request, &req); err != nil {
				if err := json.Unmarshal(request, &dummy); err == nil {
					fmt.Println("received float number" + "  request: " + string(request))
					conn.Write([]byte("{\"method\":\"isPrime\",\"prime\":false}\n"))
					continue
				}
				fmt.Println("json error: " + err.Error() + " request: " + string(request))
				conn.Write([]byte(err.Error()))
				return
			}
			if req.Method != "isPrime" {
				//fmt.Println("Method not found")
				conn.Write([]byte("Method not found"))
				return
			}
			if !strings.Contains(string(request), "number") {
				conn.Write([]byte("no number!"))
				return
			}
			inInt := req.Number
			println("correct request: " + string(request))
			//fmt.Printf("float: %f, int: %d\n", req.Number, inInt)
			if isPrime(inInt) {
				//example : {"method":"isPrime","prime":false}
				res := "{\"method\":\"isPrime\",\"prime\":true}\n"
				//fmt.Printf("request: "+string(request), len(request))
				if _, err := conn.Write([]byte(res)); err != nil {
					fmt.Println(err)
					//panic(err)
				}
			} else {
				res := "{\"method\":\"isPrime\",\"prime\":false}\n"
				//fmt.Printf("Not Prime! num: %d, res: %s\n", inInt, res)
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
