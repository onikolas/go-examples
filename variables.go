package main

import "fmt"

func main() {

	var aVar int
	anotherVar := 9.0
	fmt.Println("hi, these are the variables: ", aVar, anotherVar)

	// uncommenting any of these causes an error

	//var aVar float32
	//aVar := 5
	//aVar = anotherVar
	//AVar := 10

	var a, b, c int = 1, 3, 4
	d, e, f := "text is good ", 4, " you "

	fmt.Println(a, b, c, d, e, f, PI)
}

const PI = 3.14159265359
