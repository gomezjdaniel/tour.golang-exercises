package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	prev := 0.0

	for {
		z = z - (z*z-x)/(2*z)
		if math.Abs(z-prev) < 1e-12 {
			break
		}
		prev = z
	}

	return prev
}

func main() {
	fmt.Println("Self-made Sqrt(2): ", Sqrt(2))
	fmt.Println("     math.Sqrt(2): ", math.Sqrt(2))
}
