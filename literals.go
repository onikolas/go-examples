package main

import "fmt"

func main() {

	// 1D slice literal
	var s1 []int = []int{1, 2, 3}
	fmt.Println("1D 1:", s1)

	// equivalently
	s2 := []int{1, 2, 3}
	fmt.Println("1D 2:", s2)

	// 2D slice literal
	var d1 [][]string = [][]string{
		{"1", "2"},
		{"1", "3"},
		{"1", "1"}, //note this exta comma
	}
	fmt.Println("2D 1:", d1)

	// extra comma not needed for one liners
	var d2 [][]string = [][]string{{"1", "2"}, {"1", "3"}, {"1", "1"}}
	fmt.Println("2D 2:", d2)

	// nor in this case (ending with })
	d3 := [][]string{
		{"1", "2"},
		{"1", "3"},
		{"1", "1"}}
	fmt.Println("2D 3:", d3)

	// same but shorter (no var)
	d4 := [][]string{{"1", "2"}, {"1", "3"}, {"1", "1"}}
	fmt.Println("2D 4:", d4)

	// struct literals
	type coord struct {
		lat, long float64
	}

	a1 := coord{30.3, 80.2} // could also be: var a1 coord = coord{30.3, 80.2}
	fmt.Println("struct:", a1)

	// slice of structs
	sa1 := []coord{
		{1.1, 2.2},
		{43.3, 55.5},
	}
	fmt.Println("Slice of structs:", sa1)

	// structs containing structs
	type tracker struct {
		name     string
		location coord
	}
	fb := tracker{"bob", coord{77.7, 88.8}}
	fmt.Println("Composite struct 1:", fb)

	type coord3ple struct {
		c1, c2, c3 coord
	}
	c3 := coord3ple{coord{77.7, 88.8}, coord{7.7, 8.8}, coord{77.7, 88.8}}
	fmt.Println("Composite struct 2:", c3)

	// We can also use field names (less confusing in some cases)
	c4 := coord3ple{
		c1: coord{77.7, 88.8},
		c3: coord{1, 1},
		// c2 gets zero value
	}
	fmt.Println("Composite struct 3:", c4)

	// Declare and initialize (this is a non-reusable struct)
	onceoff := struct {
		a, b int
		c    coord
	}{
		7, 8, coord{7, 8},
	}
	fmt.Println("One-off struct:", onceoff)

	// Map literals
	mmm := map[int]string{
		1: "ena",
		2: "dyo",
		3: "tria",
	}
	fmt.Println("map: ", mmm)

	// 2D slice of maps
	pain := [][]map[int]int{
		[]map[int]int{
			map[int]int{1: 0, 2: 1},
			map[int]int{3: 2, 4: 3},
		},
		[]map[int]int{
			map[int]int{1: 0, 2: 1},
			map[int]int{3: 2, 4: 3},
		},
	}
	fmt.Print(pain)
}
