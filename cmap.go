package main

import (
	"fmt"
	"github.com/onikolas/conkmap"
)

func main() {

	amap := conkmap.New()
	amap.Set(1, "bob")
	amap.Set(2, "alice")

	fmt.Println(amap.Get(2))

	// 1. Have a look at conkmap code AND test
	// 2. Spawn some goroutines to add elements to the map concurrently
	// 3. Spawn some goroutines to read the elements and store them in a slice
	// 4. How many goroutines should we use for (2) and how many (3)? Use the time command!
}
