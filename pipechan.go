package main

import "fmt"
import "time"

func Fib(f0, f1, stop int, c chan int) {
	for i := 0; i < stop; i++ {
		f0, f1 = f1, f0+f1
		fmt.Printf("fib:%v \n", f0)
		c <- f0
	}
}

func Printer(stop int, c chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("pop:", <-c)
	}
}

func main() {

	c := make(chan int, 5)
	// c := make(chan int) // also OK, difference?
	go Fib(0, 1, 10, c)
	go Printer(10, c)

	time.Sleep(time.Second * 5)
}
