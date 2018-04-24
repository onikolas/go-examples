package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	com1 := make(chan string)
	com2 := make(chan string)

	go HttpGetter("https://en.wikipedia.org/wiki/Special:Random", com1)
	go HttpGetter("https://en.wikipedia.org/wiki/Special:Random", com2)

	select {

	case msg := <-com1:
		fmt.Println("Com1 was faster!")
		fmt.Println(msg)

	case msg := <-com2:
		fmt.Println("Com2 was faster!")
		fmt.Println(msg)

	case <-time.After(time.Second * 1):
		fmt.Println("Too slow!")

		/*default:
		fmt.Println("Too slow!")*/
	}

}

// Get a URL's body
func HttpGetter(url string, c chan string) {
	res, err := http.Get(url)
	if err != nil {
		c <- ""
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	c <- string(body[:1000])
}
