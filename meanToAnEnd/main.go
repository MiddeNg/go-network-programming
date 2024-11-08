package main

import (
	"encoding/binary"
	"fmt"
	"github.com/MiddeNg/go-network-programming/server"
	"io"
	"net"
)

func main() {
	server.StartTCPServer(meanToAnEnd)
}

type asset struct {
	Timestamp int32
	Price     int32
}
type assets struct {
	data []asset
}

func (a *assets) Exists(asset asset) bool {
	for _, v := range a.data {
		if v.Timestamp == asset.Timestamp {
			return true
		}
	}
	return false
}

func (a *assets) Add(asset asset) error {
	if a.Exists(asset) {
		return fmt.Errorf("asset already exists")
	}
	a.data = append(a.data, asset)
	return nil
}

func (a *assets) Mean(start, end int32) int32 {
	if start > end {
		return 0
	}
	var sum int64
	var count int64 = 0
	for _, v := range a.data {
		if v.Timestamp >= start && v.Timestamp <= end {
			sum += int64(v.Price)
			count += 1
		}
	}
	if count == 0 {
		return 0
	}
	return int32(sum / count)
}

func meanToAnEnd(conn net.Conn) {
	defer conn.Close()
	requestLen := 9
	buf := make([]byte, 0, requestLen)
	assets := assets{}
	for {
		for {
			n, err := conn.Read(buf[len(buf):cap(buf)])
			buf = buf[:len(buf)+n]
			if err != nil {
				if err == io.EOF {
					return
				}
				panic(err)
			}
			if len(buf) == requestLen {
				break
			}
		}
		action := fmt.Sprintf("%c", buf[0])
		if action != "Q" && action != "I" {
			print("Invalid action")
			return
		}
		timestamp := int32(binary.BigEndian.Uint32(buf[1:5]))
		price := int32(binary.BigEndian.Uint32(buf[5:9]))
		buf = buf[:0]

		fmt.Printf("Action: %s, Timestamp: %d, Price: %d\n", action, timestamp, price)
		if action == "Q" {
			mean := assets.Mean(timestamp, price)
			res := make([]byte, 4)
			binary.BigEndian.PutUint32(res, uint32(mean))
			conn.Write(res)
			continue
		}
		if err := assets.Add(asset{Timestamp: timestamp, Price: price}); err != nil {
			conn.Write([]byte("timestamp already exists"))
			return
		}

	}
}
