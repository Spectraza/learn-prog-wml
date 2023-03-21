package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

	for {
		ran := rand.Intn(11)
		time.Sleep(1 * time.Second)
		fmt.Println(ran)
		if ran == 10 {
			break
		}
	}
}
