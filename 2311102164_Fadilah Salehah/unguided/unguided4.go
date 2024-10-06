package main

import (
	"fmt"
)

func main() {
	var beratParselInGram, totalBiaya float64
	var beratKg, beratGram int

	fmt.Print("Masukkan berat parsel (dalam gram): ")
	fmt.Scan(&beratParselInGram)

	beratKg = int(beratParselInGram) / 1000
	beratGram = int(beratParselInGram) % 1000

	totalBiaya = float64(beratKg) * 10000

	if beratGram >= 500 {
		totalBiaya += 2500
	} else if beratGram > 0 {
		totalBiaya += float64(beratGram) * 15
	}

	if beratKg >= 10 {
		totalBiaya = 100000
	}

	fmt.Printf("\nBerat parsel (gram): %.0f\n", beratParselInGram)
	fmt.Printf("Detail berat: %d kg + %d gram\n", beratKg, beratGram)
	fmt.Printf("Total biaya: Rp. %.0f\n", totalBiaya)
}
