package main

import (
	"fmt"
	"os"
)

type Storer interface {
	Store(string)
	Retrieve(int) string
}

type FancySlice struct {
	s []string
}

// by implementing a Store and a Retrive method FancySlice implements Storer
func (f *FancySlice) Store(a string) {
	f.s = append(f.s, a)
}

func (f FancySlice) Retrieve(i int) string {
	if i < 0 || i >= len(f.s) {
		return ""
	}
	return f.s[i]
}

// A type can implement multiple interfaces. This is a common from the standard library:
//type Stringer interface {
//    String() string
//}
func (f FancySlice) String() string {
	var s string
	for _, v := range f.s {
		s = s + v + ","
	}
	return s
}

func main() {
	var storage Storer
	storage = &FancySlice{}

	storage.Store("A string")
	storage.Store("Another string")
	fmt.Println(storage.Retrieve(0))
	// ...what did we gain over declaring storage as:
	// var storage FancySlice{}

	// stringer doing work
	fmt.Println(storage)

	// save to disk
	SaveToFile(storage, "abc.txt")
}

// TODO: implement a map storer

// Can save *any* Storer to disk ...with some help
// Add another method to Storer to allow SaveTofile to work correctly
func SaveToFile(s Storer, filename string) {

	//ignoring the error for brievity
	file, _ := os.Create(filename)
	defer file.Close()

	data := s.Retrieve(0)
	bytes := []byte(data)

	file.Write(bytes)
}
