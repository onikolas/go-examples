package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {

	// ifs can have initializers
	if a := 10; a > 9 {
		fmt.Println("Trivial is trivial")
	}

	if val, err := RiskyFunc(9, 0); err != nil {
		fmt.Println(val)
		panic("oh noes!")
	}

	// whats wrong with this ?
	// fmt.Println(val + 1)

	// reader gets stuff from the console (os.stdin)

	num := rand.Intn(10)
	switch num {
	case 0:
		fmt.Println("Unlucky")
	case 7:
		fmt.Println("Lucky!")
	default:
		fmt.Println("Meh..")
	}
}

func RiskyFunc(a, b int) (int, error) {
	if b != 0 {
		return a / b, nil
	}
	return 0, errors.New("division by zero")
}
