//go:build range
// +build range

package main

import "fmt"

func main() {
	pow := make([]int, 10)
	fmt.Println(pow)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for i, value := range pow {
		fmt.Printf("%d %d\n", i, value)
	}
}
