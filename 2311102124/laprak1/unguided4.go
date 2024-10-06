/*Liya Khoirunnisa - 2311102124*/

package main

import (
	"fmt" // Untuk input atau output
)

func main() {
	// Input berat parsel dalam gram
	var beratGram int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&beratGram)

	// Menghitung berat total dalam kg dan gram
	beratKg := beratGram / 1000
	sisaGram := beratGram % 1000

	// Menghitung biaya kirim
	biayaPerKg := 10000
	biayaSisa := 0

	// Cek 
	if beratKg > 10 {
		biayaSisa = 0
	} else if sisaGram >= 500 {
		biayaSisa = sisaGram * 5
	} else {
		biayaSisa = sisaGram * 15
	}

	// Menghitung total biaya kirim
	totalBiaya := (beratKg * biayaPerKg) + biayaSisa

	// Output detail berat dan biaya
	fmt.Printf("Detail berat: %d kg + %d gr\n", beratKg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", beratKg*biayaPerKg, biayaSisa)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
