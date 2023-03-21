package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	num := 28
	fmt.Println(num)

	for num > 10 {
		num /= 2
	}
	fmt.Println(num)
}
