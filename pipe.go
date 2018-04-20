package main

import (
	"fmt"
	"sync"
	"time"
)

// FIFO
type ConkQ struct {
	data       []int
	head, tail int
	lock       sync.Mutex
}

func (cq *ConkQ) Push(a int) {
	cq.lock.Lock()
	defer cq.lock.Unlock()
	next := (cq.head + 1) % len(cq.data)

	// Q full
	for cq.tail == next {
		fmt.Println("full!")
		cq.lock.Unlock()
		time.Sleep(time.Millisecond * 500)
		cq.lock.Lock()
	}

	cq.head = next
	cq.data[next] = a
}

func (cq *ConkQ) Pop() int {
	cq.lock.Lock()
	defer cq.lock.Unlock()

	// Q empty
	for cq.head == cq.tail {
		fmt.Println("empty!")
		cq.lock.Unlock()
		time.Sleep(time.Millisecond * 500)
		cq.lock.Lock()
	}

	cq.tail = (cq.tail + 1) % len(cq.data)
	data := cq.data[cq.tail]
	return data
}

func Fib(f0, f1, stop int, cq *ConkQ) {
	for i := 0; i < stop; i++ {
		f0, f1 = f1, f0+f1
		fmt.Printf("fib:%v \n", f0)
		cq.Push(f0)
	}
}

func Printer(stop int, cq *ConkQ) {
	for i := 0; i < 10; i++ {
		fmt.Println("pop:", cq.Pop())
	}
}

func main() {

	cq := ConkQ{
		data: make([]int, 15),
	}

	// can we swap these ?
	go Fib(0, 1, 10, &cq)
	go Printer(10, &cq)

	time.Sleep(time.Second * 5) //why?
}
