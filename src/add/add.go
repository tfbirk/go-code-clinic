package main

import "fmt"

func add (x int, y int) int {
	return x + y
}

func main () {
	fmt.Println("42 + 13 = ", add(42, 13))
}