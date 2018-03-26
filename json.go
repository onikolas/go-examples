package main

import (
	"encoding/json"
	"fmt"
)

type Playlist struct {
	Songs []string
	Fav   int
}

func main() {

	var myQueenList Playlist
	myQueenList.Songs = append(myQueenList.Songs, "Bohemian rhapsody")
	myQueenList.Songs = append(myQueenList.Songs, "Another one bites the dust")
	myQueenList.Songs = append(myQueenList.Songs, "Under pressure")
	myQueenList.Fav = 1

	LedZepList := Playlist{
		[]string{"Stairway", "My little pony"},
		1,
	}

	// Problems ?

	// Emulating json:
	fmt.Println(myQueenList)
	fmt.Println(LedZepList)

	// Actually json:
	bytes, _ := json.Marshal(myQueenList)
	fmt.Println(string(bytes))
	bytes, _ = json.Marshal(LedZepList)
	fmt.Println(string(bytes))

}
