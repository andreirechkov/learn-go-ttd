package main

import (
	"fmt"
	"math"
)

const epsilon = 1e-9

func areFloatsEqual(a, b float64) bool {
    return math.Abs(a - b) < epsilon
}


func Sqrt(x float64) float64 {
	z := x / 2
	prev_value := z

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		
		if areFloatsEqual(prev_value, z) {
			break
		}
		
		fmt.Printf("Iter %d: z = %f\n", i+1, z)
		prev_value = z
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
