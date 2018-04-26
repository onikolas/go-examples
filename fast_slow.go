package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Sender(n, speed int) chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < n; i++ {
			c <- rand.Intn(10)
			time.Sleep(time.Millisecond * time.Duration(speed))
		}
		close(c)
	}()

	return c
}

func SpawnSenders(n int) []chan int {
	//slice of chans
	cs := make([]chan int, n)

	for i := 0; i < len(cs); i++ {
		if i%2 == 0 {
			cs[i] = Sender(10, 100) //fast sender
		} else {
			cs[i] = Sender(10, 1000) //slow sender
		}
	}

	return cs
}

func FanIn(c []chan int) chan int {
	cs := make(chan int)
	done := make(chan bool)
	for i := 0; i < len(c); i++ {
		go func(a chan int) {
			for j := range a {
				cs <- j
			}
			done <- true
		}(c[i])
	}

	go func() {
		for i := 0; i < len(c); i++ {
			<-done
		}
		close(cs)
	}()
	return cs
}

func Receiver(c chan int) {
	count := 0
	for i := range c {
		fmt.Println(i)
		count++
	}
	fmt.Println("Got ", count)
}

func main() {
	Receiver(FanIn(SpawnSenders(10)))
}
