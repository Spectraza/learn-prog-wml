package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	a := rand.Intn(100)
	b := math.Sqrt(float64(a))
	fmt.Println(b)
}
