package main

import "fmt"

func main() {
	defer fmt.Println("again")
	defer fmt.Println("world")
	fmt.Println("hello")

}
