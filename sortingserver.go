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

/* var data = []int{8, 9, 122, 1, 1, 100, 999, 1001, 0} */

/* [ [1,2,3,4,5],
   [1,3,3,4,4],
   [1,1,1,1,1] ]
*/

package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	datastore, _ := ReadFile("sortingexample.csv")
	cleanData := ConvertToSlice(datastore)
	sortedArray := CountSort(cleanData)
	fmt.Println(datastore)
	fmt.Println(cleanData)
	fmt.Println(sortedArray)
}

//ReadFile takes a file as input and writes the content of the file as a 2-d array of strings
func ReadFile(f string) ([][]string, error) {
	openFile, err := os.Open(f)
	if err != nil {
		panic("file could not be read")
	}
	defer openFile.Close()

	g := csv.NewReader(openFile)

	data, err := g.ReadAll()
	if err != nil {
		return nil, errors.New("failed to read file")
	}

	return data, nil
}

// TODO Right now, the ConvertToSlice function fails if there are spaces between values in the csv, or there are commas at the end of the line. Make the implementation tolerate these faults

func ConvertToSlice(inputData [][]string) []int {
	var destination []int
	for i := range inputData {
		for j := range inputData[i] {
			v := inputData[i][j]
			w, err := strconv.Atoi(v)
			if err != nil {
				log.Printf("Input %v , %v from the input file is not an integer", i, j)
			}
			destination = append(destination, w)
		}
	}
	return destination
}

// CountSort sorts an integer slice (https://en.wikipedia.org/wiki/Counting_sort)
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


