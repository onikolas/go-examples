package main

import (
	"fmt"
	"math/rand"
	//	"time"
)

func Sender(n, speed int) chan int {
	c := make(chan int, 2)

	go func() {
		for i := 0; i < n; i++ {
			c <- rand.Intn(10)
			//time.Sleep(time.Millisecond * time.Duration(speed))
		}
		close(c) //try commenting out
	}()

	return c
}

func Receiver(c chan int) {
	for i := range c {
		fmt.Println("Got ", i)
	}
}

func main() {

	Receiver(Sender(10, 100))
}
