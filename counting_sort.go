// Implement counting sort.
// This is a warmup so make it as simple as possible:
// * Data are integers
// * Data declared inside the source code
// * Hard-coded index limit
// * Plain, boring sequential program

package main

import "fmt"

var data = []int{8, 9, 122, 1, 1, 100, 999, 1001, 0}

func CountSort(input []int) []int {
	sorted := make([]int, len(input))
	size := input[0]
	for i := range input {
		if input[i] > size {
			size = input[i]
		}
	}

	auxArray := make([]int, size)
	for i := range input {
		auxArray[input[i]]++
	}

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
