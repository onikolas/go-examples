package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

func Prime(n int) bool {

	if n <= 1 {
		return false
	} else if n <= 3 {
		return true
	} else if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i = i + 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}

func MarkPrimes(data []int, done chan bool) {
	for i := 0; i < len(data); i++ {
		if Prime(data[i]) {
			data[i] = 1
		} else {
			data[i] = 0
		}
	}

	// signal that we are finished
	done <- true
}

func main() {

	data := make([]int, 1024*1024)

	for i := 0; i < len(data); i++ {
		data[i] = rand.Intn(6000) //be carefull of this number!
	}

	fmt.Println("using", runtime.NumCPU(), " goroutines")

	chunkSize := len(data) / runtime.NumCPU()

	done := make(chan bool)

	for i := 0; i < runtime.NumCPU(); i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize //problem?
		go MarkPrimes(data[start:end], done)
	}

	// sync
	for i := 0; i < runtime.NumCPU(); i++ {
		<-done
	}

	count := 0
	for i := 0; i < len(data); i++ {
		if data[i] == 1 {
			count++
		}
	}

	fmt.Println(count, " primes")

	// Task: Start counting as soon as a chunk is finished.
}
