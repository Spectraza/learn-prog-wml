package main

import "math"

func integrate(f func(float64) float64, a, b float64) float64 {
	integral := 0.0
	dx := 0.001
	for x := a; x < b; x += dx {
		value := f(x)
		s := value * dx
		integral += s
	}
	return integral
}

func main() {
	println(integrate(math.Sqrt, 0, 10))
}
