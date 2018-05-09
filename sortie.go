// Implement sortie.

// Sortie accepts data from various sources and sorts them. On request, it outputs the sorted data to
// either the screen or a file. We have outside guarantees that our data will be dense so counting sort
// makes sense as the underlying algorithm. Sortie must be able to sort any data that implements the
// Sorter interface (we must define what this is).
//
// Can accept data from:
// -A file, given at program invocation.
// -The network, on a port given as command line argument
// -Stretch goal: Monitor the file after loading it (assuming it only increases in size - e.g. log)
//
// The request to display/save the sorted data can be a key press, a net package or both

package main

import (
	"encoding/csv"
	"fmt"
	"log"
)

func main() {

	var data [][]string = [][]string{{"1", "2", "3", "48", "5"}, {"1", "3", "3", "4", "4"}, {"1", "1", "1", "1", "1"}}
	x := ConvertToSlice(data)
	fmt.Println(x)
}

func ConvertToSlice(inputData [][]string) []int {
	var destination []int
	for i := range inputFile {
		for j := range inputFile[i] {
			v := inputFile[i][j]
			w, err := strconv.Atoi(v)
			if err != nil {
				log.Printf("Input %v , %v from the input file not an integer", i, j)
			}
			destination = append(destination, w)
		}
	}
	return destination
}
