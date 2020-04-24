package main

import "fmt"

type Inter interface {
	Do()
}

type A struct {
	Inter
	Name string
	// B
}

type B struct {
	Name string
}

// Do something
func (b B) Do() {
	b.Name = "B.name"
	fmt.Printf("%+v\n", b.Name) // output for debug

}

// // Do ...
// func (a A) Do() {
// 	fmt.Printf("%+v\n", "A.Do") // output for debug
// }

func main() {
	var a = &A{Inter: B{}}
	a.Do()
}
