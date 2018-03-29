package main

import "fmt"

func main() {

	var a int = 10
	var b *int = &a
	c := &a
	fmt.Println(a, b, *b, c, *c)

	a = 11
	fmt.Println(a, b, *b, c, *c)

	d := c
	fmt.Println(*d)

	*d = *d + 1
	fmt.Println(a, *d)
	// what would this do ?
	// d = d + 1

	SpaghettiDoubler(&a)
	fmt.Println(a, b, *b, c, *c)

	a = FunctionalDoubler(a)
	fmt.Println(a, b, *b, c, *c)

	array := [3]int{4, 8, 16}
	ArrayDoubler(array)
	fmt.Println("array: ", array)

	slice := array[:]
	SliceDoubler(slice)
	fmt.Println("slice: ", slice)
	// what's the problem with this:
	//ArrayDoubler(slice)

	var T struct {
		a, b int
	}
	T.b = 4
	fmt.Println(T)

	pT := &T
	pT.a = 6
	fmt.Println(T, pT)
}

// bad idea
func SpaghettiDoubler(a *int) {
	*a = *a * 2
}

// good idea
func FunctionalDoubler(a int) int {
	return a * 2
}

// use case ?
func ArrayDoubler(a [3]int) {
	for i := 0; i < 3; i++ {
		a[i] = a[i] * 2
	}
}

// bad idea ?
func SliceDoubler(a []int) {
	for i := 0; i < 3; i++ {
		a[i] = a[i] * 2
	}
}
