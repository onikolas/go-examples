package main

import (
	"fmt"
)

func main() {

	i := 0
	for i < 10 {
		fmt.Print(i, " ")
		i = i + 1
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i, " ")
	}

	for {
		fmt.Println("Control+C ! ")
	}
}

// Write for loops to produce the following outputs:

// ******
// ******
// ******
// ******
// ******

// *
// **
// ***
// ****
// *****
// ******
// *******

// ******
// ****
// **

// ***
// **
// *
// **
// ***

// 1
// 12
// 1234
// 12345678

// ******
// *    *
// *    *
// *    *
// ******
