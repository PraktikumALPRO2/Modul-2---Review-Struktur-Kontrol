package main

import (
	"fmt"
	"math"
)

func f(k int) float64 {
	return math.Pow(float64(4*k+2), 2) / float64((4*k+1)*(4*k+3))
}

func sqrt2Approximation(k int) float64 {
	approx := 0.0
	for i := 0; i <= k; i++ {
		approx += f(i)
	}
	return approx
}

func main() {
	var k int
	for {
		fmt.Print("Nilai K = ")
		_, err := fmt.Scan(&k)
		if err != nil {
			break
		}
		approx := sqrt2Approximation(k)
		fmt.Printf("Nilai akar 2 %.10f\n", approx)
	}
}
