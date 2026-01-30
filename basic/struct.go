//go:build structs
// +build structs

package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	strt := Vertex{1, 2}
	fmt.Println(strt)
}
