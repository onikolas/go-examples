package main

import "fmt"

func main() {

	// uncomment and run each section one at a time using this command:
	// time go run slices.go

	a := []int{}
	for i := 0; i < 2048*2048*16; i++ {
		a = append(a, i)
	}

	/*a := make([]int, 2048*2048*16)
	for i := 0; i < 2048*2048*16; i++ {
		a[i] = i
	}*/

	/*a := make([]int, 0, 2048*2048*16)
	for i := 0; i < 2048*2048*16; i++ {
		a[i] = i
	} // tip: print a's len() and cap() */

	/*a := make([]int, 0, 2048*2048*16)
	for i := 0; i < 2048*2048*16; i++ {
		a = append(a, i)
	}*/

	fmt.Println(a[len(a)-1])

	//ShowSlice([]int{1, 2, 3, 4})
}

func ShowSlice(a []int) {

	for i := 0; i < len(a); i++ {
		fmt.Printf("element %v has value %v \n", i, a[i])
	}

	for i, v := range a {
		fmt.Printf("element %v has value %v \n", i, v)
	}
}
