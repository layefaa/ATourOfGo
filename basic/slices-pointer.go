//go:build slices_pointer
// +build slices_pointer

package main

import "fmt"

func main() {
	names := []string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names)

	a := names[0:1]
	b := names[1:4]

	fmt.Println(a, b)

	b[0] = "XXXX"

	fmt.Println(names)
	fmt.Println(a, b)

}
