//go:build slice_literals
// +build slice_literals

package main

import "fmt"

func main() {
	q := []int{2, 3, 5}
	fmt.Println(q)

	r := []bool{true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
	}
	fmt.Println(s)

	z := [8]int{
		1, 2, 3, 4, 5, 6, 7, 8,
	}

	a := z[:8]
	b := z[4:]

	fmt.Println(a, b)

}
