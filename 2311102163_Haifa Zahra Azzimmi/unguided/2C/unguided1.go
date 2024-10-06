package main

import (
	"fmt"
)

func main() {
	var beratParsel, totalBiaya float64
	var kg, gram int

	fmt.Print("Masukkan berat parsel (dalam gram): ")
	fmt.Scan(&beratParsel)

	kg = int(beratParsel) / 1000
	gram = int(beratParsel) % 1000

	totalBiaya = float64(kg) * 10000

	if gram >= 500 {
		totalBiaya += 2500
	} else if gram > 0 {
		totalBiaya += float64(gram) * 15
	}

	if kg >= 10 {
		totalBiaya = 100000
	}

	fmt.Printf("\nBerat parsel (gram): %.0f\n", beratParsel)
	fmt.Printf("Detail berat: %d kg + %d gram\n", kg, gram)
	fmt.Printf("Total biaya: Rp. %.0f\n", totalBiaya)
}
