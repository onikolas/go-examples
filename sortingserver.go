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
"fmt"
"log"
// "encoding/csv"
"strconv"
)

/* var data = []int{8, 9, 122, 1, 1, 100, 999, 1001, 0} */

/* [ [1,2,3,4,5],
   [1,3,3,4,4],
   [1,1,1,1,1] ]
*/


package main

import (
"encoding/csv"
"fmt"
"log"
)

func main() {

	data := [][]string{{"1", "2", "3", "48", "5"}, {"ASD", "3", "3", "4", "4"}, {"1", "1", "1", "1", "1"}}
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
				log.Printf("Input %v , %v from the input file is not an integer", i, j)
			}
			destination = append(destination, w)
		}
	}
	return destination
}
/*

// CounSort sorts an integer slice (https://en.wikipedia.org/wiki/Counting_sort)
func CountSort(input []int) []int {
  sorted := make([]int, len(input))
  size := input[0]

  // find the size
  for i := range input {
    if input[i] > size {
      size = input[i]
    }
  }

  // populate index array
  auxArray := make([]int, size+1)
  for i := range input {
    auxArray[input[i]]++
  }

  // save the sorted results
  count := 0
  for i := range auxArray {
    for j := 0; j < auxArray[i]; j++ {
      sorted[count] = i
      count++
    }
  }

  return sorted
}

func main() {
  results := CountSort(data)
  fmt.PrintLn(results)
}