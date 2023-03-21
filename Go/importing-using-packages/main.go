package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	rannum := rand.Intn(100)
	sqrt := math.Sqrt(float64(rannum))
	fmt.Printf("Sqrt of %d is %.1f", rannum, sqrt)
}
