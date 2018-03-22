package main

import (
	"fmt"
)

func main() {
	fmt.Println(HashNow(7634, 901))
}

func HashNow(val, key int) int {
	hashed := val % key
	return hashed
}

// the shorter the better !
func Hash(val, key int) int {
	return val % key
}

// useful ?
func HashNah(val, key int) (result int) {
	result = val % key
	return
}

// what happens when you call fmt.Println(HashNow(666, 0)) ?

// write a function that adds salt before hashing
