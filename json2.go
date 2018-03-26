package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Playlist struct {
	Songs []string
	Fav   int
}

// Paste this in a file called list.txt
// {"Songs":["Stairway","My little pony"],"Fav":1}

func main() {
	dat, _ := ioutil.ReadFile("list.txt")
	fmt.Println(dat)
	// Assume no JSON, what now ?

	var list Playlist
	json.Unmarshal(dat, &list)

	fmt.Println(list)
}
