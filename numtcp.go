package main

import (
	"net"
	"fmt"
	"math/rand"
	"time"
)


func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:34567")

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	for i := 0; i < 10; i++ {
		fmt.Fprintf(conn, "%d ", rand.Intn(30))
		time.Sleep(time.Millisecond * 1000)
	}
}