package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":34567")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		fmt.Println(conn.LocalAddr().String(), " has connected")
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {

	reader := bufio.NewReader(conn)

	for {
		// TODO add read deadline
		str, err := reader.ReadString(' ')
		if err != nil {
			fmt.Println(conn.LocalAddr().String(), " has ", err)
			conn.Close()
			return
		} else {
			fmt.Println("Read ", str)
		}
	}
}
