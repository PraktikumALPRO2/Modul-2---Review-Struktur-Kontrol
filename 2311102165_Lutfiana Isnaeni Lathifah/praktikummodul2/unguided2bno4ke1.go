// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import (
	"fmt"
)

func f(k int) float64 {
	// Menghitung nilai f(k) = ((4k + 2)^2) / ((4k + 1)(4k + 3))
	numerator := (4*float64(k) + 2) * (4*float64(k) + 2)
	denominator := (4*float64(k) + 1) * (4*float64(k) + 3)
	return numerator / denominator
}

func main() {
	var k int

	// Meminta input nilai K
	fmt.Print("Nilai K = ")
	fmt.Scan(&k)

	// Menghitung dan menampilkan nilai f(K)
	result := f(k)
	fmt.Printf("Nilai f(K) = %.10f\n", result)
}
