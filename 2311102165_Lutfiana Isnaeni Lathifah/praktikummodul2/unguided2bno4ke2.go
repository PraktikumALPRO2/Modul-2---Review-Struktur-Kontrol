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
	var K int
	var sqrt2 float64 = 1

	// Meminta input nilai K
	fmt.Print("Nilai K = ")
	fmt.Scan(&K)

	// Menghitung hampiran akar 2
	for k := 0; k <= K; k++ {
		sqrt2 *= f(k)
	}

	// Menampilkan hasil hampiran akar 2
	fmt.Printf("Nilai akar 2 = %.10f\n", sqrt2)
}
