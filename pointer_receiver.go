package main

import (
	"fmt"
)

type Face interface {
	Smile()
	Scowl()
}

type Happy string

func (h Happy) Smile() {
	fmt.Println(":)")
}

func (h *Happy) Scowl() {
	fmt.Println(":(")
}

func EmotionProcessor(face Face) {
	face.Smile()
	face.Scowl()
}

func main() {
	var a Happy
	EmotionProcessor(a) // error: Happy does not implement Face, only *Happy does
	EmotionProcessor(&a)
}
