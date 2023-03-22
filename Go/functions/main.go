package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func divine(x, y float64) float64 {
	return x / y
}

func main() {
	fmt.Println(add(2, 2), divine(24, 2))
}
