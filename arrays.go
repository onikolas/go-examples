package main

import (
	"fmt"
	//	"math/rand"
)

// first see https://tour.golang.org/moretypes/6 - 9

func main() {

	// whats the difference between
	a := [3]int{1, 2, 3}
	// and
	b := []int{1, 2, 3}

	fmt.Println(a, b)

	// what happened to {1,2,3} ?
	b = []int{2, 3, 4, 5}
	fmt.Println(b)

	// slicing
	fmt.Println(b[0:2], b[2:4])
	// slicing with defaults
	fmt.Println(b[:2], b[2:])
	// use case?
	fmt.Println(len(b), cap(b), len(b[:2]), cap(b[:2]))

	// dynamic arrays
	c := make([]float32, 16)
	fmt.Println(c)

	// populate c using a for loop and rand.Float32()

	// what if c was 1GB ? how do we populate faster ?
}
