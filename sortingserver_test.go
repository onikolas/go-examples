package main

import "testing"

func TestSortFile(T *testing.T) {
	type result struct {
		filename string
		result []int
	}

	r := result{
		"sortingexample.csv",
		[]int{1, 1, 1, 1, 2, 3, 3, 3, 4, 4, 5, 48, 126, 234, 435},
	}

	datastore, _ := ReadFile(r.filename)
	cleanData := ConvertToSlice(datastore)
	sortedArray := CountSort(cleanData)

	if len(sortedArray) != len(r.result) {
		T.Error("encountered different array sizes")
	}
	for i := 0; i < len(sortedArray); i++ {
		if sortedArray[i] != r.result[i] {
			T.Errorf("sorting result different than expected, expected %d, got %d", r.result[i], sortedArray[i])
		}
	}
}