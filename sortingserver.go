// Sorting server accepts data from various sources and sorts them. On request, it outputs the sorted
// data to  either the screen or a file. We have outside guarantees that our data will be dense so
// counting sort  makes sense as the underlying algorithm. Sortie must be able to sort any data
// that implements the  Sorter interface (we must define what this is).
//
// Can accept data from:
// -A file, given at program invocation.
// -The network, on a port given as command line argument
// -Stretch goal: Monitor the file after loading it (assuming it only increases in size - e.g. log)
//
// The request to display/save the sorted data can be a key press, a net package or both

package main

import (
	/*
	"encoding/csv"

	"fmt"
	"log"
	"os"

	"time"
	*/
	"sync"
	"fmt"
	"time"
	"math/rand"
	"bufio"
	"os"
	"net"
	"strconv"
	"strings"
)

// This interface describes all the data types that can be sorted with this program. All sortable data types can be assigned an index for sorting.
type Sorter interface {
	indexOf() int
}

type sortablestring string

func(ss sortablestring) indexOf() int {
	return int(ss[0])
}

type sortableinteger int

func(si sortableinteger) indexOf() int {
	return int(si)
}

// Stores data to be sorted with a mutex for concurrent go-routine access
type DataSort struct{
	rawData []int
	locker sync.RWMutex
}

// Initialises a DataSort structure with the requisite length. largest represents the largest element we can get.
func NewDataSort(largest int) DataSort {
	var holder DataSort
	holder.rawData = make([]int, largest)
	return holder
}

//
func (ds *DataSort) Add(a ...int) {
	ds.locker.Lock()
	for _, num := range a {
		ds.rawData[num]++
	}
	ds.locker.Unlock()
}

func (ds *DataSort) Sort() []int {
	ds.locker.RLock()
	counter := len(ds.rawData)
	var results []int
	for i := 0; i < counter; i++ {
		NumDuplications := ds.rawData[i]
		for j := 0; j < NumDuplications; j++ {
			results = append(results, i)
		}
	}
	ds.locker.RUnlock()
	return results
}

// Prints a number 'quantity' of random integers
func NetworkSimulator(data *DataSort, quantity int) {
	// k represents the largest integer
	k := len(data.rawData)
	fmt.Println(k)
	for i := 0; i < quantity; i++ {
		j := rand.Intn(k)
		data.Add(j)
		time.Sleep(100*time.Millisecond)
	}
}

func ContinuousSorter(data *DataSort) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("entercommand")
		input, _ := reader.ReadString('\n')
		fmt.Println(input)
		if input == "print\n" {
			sortedArray := data.Sort()
			fmt.Println(sortedArray)
		}
	}
}

//Stores data as an array of DataSort struct in order to enable concurrency
type ArrayDataSort struct {
	segments []DataSort
}

func(ads *ArrayDataSort) AddSegment(numSegment int, segmentSize int) {
	holder := make([]DataSort, numSegment)
	for i := range holder {
		holder[i] = NewDataSort(segmentSize)
	}
	ads.segments = append(ads.segments, holder...)
}

// Initializer of ArrayDataSort
func NewArrayDataSort(segmentSize int, numberSegment int) ArrayDataSort {
	var holder ArrayDataSort
	holder.segments = make([]DataSort, numberSegment)
	for i := 0; i <numberSegment; i++ {
		holder.segments[i] = NewDataSort(segmentSize)
	}
	return holder
}


func (ds *ArrayDataSort) Add(a ...Sorter) {
	numberSegment := len(ds.segments)
	segmentLength := len(ds.segments[0].rawData)
	for _, num := range a {
		quotient := num.indexOf() / segmentLength
		remainder := num.indexOf() % segmentLength
		if num.indexOf() < 0 {
			fmt.Printf("Your supplied number is out of the range of numbers program can handle. It can between 0 and %v", numberSegment*segmentLength)
			continue
		}
		if quotient >= numberSegment {
			ds.AddSegment(quotient - numberSegment + 1, segmentLength)
			numberSegment = len(ds.segments)
		}
		ds.segments[quotient].Add(remainder)
	}
}

func (ds *ArrayDataSort) Sort() []int {
	segmentLength := len(ds.segments[0].rawData)
	var results []int
	for i := 0; i < len(ds.segments); i++ {
		storer := ds.segments[i].Sort()
		for _,num := range storer {
			results = append(results, i * segmentLength + num)
		}
	}
	return results
}

func handleConnection(conn net.Conn, Data *ArrayDataSort) {

	reader := bufio.NewReader(conn)

	for {
		// TODO add read deadline
		str, err := reader.ReadString(' ')
		if err != nil {
			fmt.Println(conn.LocalAddr().String(), " has ", err)
			conn.Close()
			fmt.Println(Data.Sort())
			return
		} else {
			intbuffer, error := strconv.Atoi(strings.TrimSpace(str))
			intbuffer1 := sortableinteger(intbuffer)
			if error != nil {
				fmt.Printf("Received data could not be converted to integer. Received (%v)", str)
			}
			Data.Add(intbuffer1)
		}
	}

}

func main() {
	MasterData := NewArrayDataSort(5,5)
	ln, err := net.Listen("tcp", ":34567")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		fmt.Println(MasterData.Sort())
		if err != nil {
			// handle error
		}
		fmt.Println(conn.LocalAddr().String(), " has connected")
		go handleConnection(conn, &MasterData)
	}


}

/*

func FileSorter(f string) []int {
	datastore, _ := ReadFile(f)
	cleanData := DataCleaner(ConvertToSlice(datastore))
	sortedArray := CountSort(cleanData)
	fmt.Println(sortedArray)
	return sortedArray
}

func ContinuousSorter(f string) []int {

}

// DataCleaner takes a string slice as input, converts all constituent strings to integers and returns an integer slice
func DataCleaner(inputData []string) []int {
	var destination []int
	for i := range inputData {
		v := inputData[i]
		w, err := strconv.Atoi(v)
		if err != nil {
			log.Printf("Input %v from the input file is not an integer", i)
		}
		destination = append(destination, w)
	}
	return destination
}

// ReadFile takes a file as input and writes the content of the file as a 2-d array of strings
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

func ConvertToSlice(inputData [][]string) []string {
	var destination []string
	for i := range inputData {
		for j := range inputData[i] {
			v := inputData[i][j]
			destination = append(destination, v)
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

*/