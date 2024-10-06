// Agnes Refilina Fiska
// S1 IF11-05 / 2311102126

package main

import (
	"fmt"
)

func main() {
	var beratParsel int

	// Input berat parsel
	fmt.Print("Masukkan berat parsel (dalam gram): ")
	_, err := fmt.Scan(&beratParsel)
	if err != nil || beratParsel < 0 {
		fmt.Println("Input tidak valid. Harap masukkan angka positif.")
		return
	}

	// Hitung berat parsel dalam kilogram dan sisa gram
	beratParselKg := beratParsel / 1000
	sisaGram := beratParsel % 1000

	// Hitung biaya dasar
	biayaDasar := 10000 * beratParselKg
	biayaTambahan := 0

	// Hitung biaya sisa gram
	if beratParselKg > 10 {
		sisaGram = 0 // Jika berat lebih dari 10 kg, sisa gram tidak dikenakan biaya
	} else if sisaGram >= 500 {
		biayaTambahan = 5 * sisaGram // Rp 5 per gram jika sisa >= 500 gram
	} else {
		biayaTambahan = 15 * sisaGram // Rp 15 per gram jika sisa < 500 gram
	}

	// Hitung total biaya
	totalBiaya := biayaDasar + biayaTambahan

	// Output
	fmt.Printf("Detail berat: %d kg + %d gram\n", beratParselKg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaDasar, biayaTambahan)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
