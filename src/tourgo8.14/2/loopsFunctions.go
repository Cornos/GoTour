package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for z != math.Sqrt(x) {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("%g\n", z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
