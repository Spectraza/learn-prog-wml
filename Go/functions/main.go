package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}

func greet(name string) (greeting string) {
	greeting = fmt.Sprintf("Hello %s \n", name)
	return
}

func main() {
	
}
