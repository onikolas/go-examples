package main

import "fmt"

type Vector3 struct {
	x, y, z float32
}

func AddVectors(a, b Vector3) Vector3 {
	return Vector3{a.x + b.x, a.y + b.y, a.z + b.z}
}

func (a Vector3) Add(b Vector3) Vector3 {
	return Vector3{a.x + b.x, a.y + b.y, a.z + b.z}
}

func main() {

	fmt.Println(AddVectors(Vector3{1, 2, 3}, Vector3{1, 2, 3}))

	a := Vector3{1, 2, 3}
	b := Vector3{3, 2, 1}

	fmt.Println(a.Add(b))

	fmt.Println(a)
	a.Normalize()
	fmt.Println(a)
}

// what happens when the func signature is this:
// func (a Vector3) Normalize() {
func (a *Vector3) Normalize() {
	m := a.x*a.x + a.y*a.y + a.z*a.z
	a.x, a.y, a.z = a.x/m, a.y/m, a.z/m
}

// Excersise: Define a data structure for storing the text of a silly text editor
// which only allows inserting and deleting whole lines. Implement the insertion and
// deletion functions as member functions.

// Before starting, why can't it be a []string ?
