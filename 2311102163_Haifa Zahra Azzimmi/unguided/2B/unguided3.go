package main

import (
	"fmt"
	"math"
)

func hitungFK(K float64) float64 {
	atas := (4*K + 2) * (4*K + 2)
	bawah := (4*K + 1) * (4*K + 3)
	return atas / bawah
}

func hitungAkar2() float64 {
	return math.Sqrt(2)
}

func main() {
	var K float64
	fmt.Print("Masukkan nilai K: ")
	fmt.Scan(&K)

	nilaiFK := hitungFK(K)
	fmt.Printf("Nilai f(K) = %.10f\n", nilaiFK)

	fmt.Println("\nMenghitung akar 2 dengan presisi hingga 10 digit desimal:")
	akar2 := hitungAkar2()
	fmt.Printf("Nilai akar 2 = %.10f\n", akar2)

	fmt.Println("\nHasil iterasi beberapa nilai K:")
	for i := 1; i <= 3; i++ {
		nilaiFK = hitungFK(float64(i))
		fmt.Printf("Nilai K = %d, Nilai f(K) = %.10f\n", i, nilaiFK)
	}
}
