package main

import (
	"fmt"
)

func main() {
	var beratParsel int
	var kg, sisaBerat int
	var biayaPerKg, biayaSisa int

	// Meminta input berat parsel dalam gram
	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&beratParsel)

	// Menghitung berat dalam kg dan sisa berat 
	kg = beratParsel / 1000
	sisaBerat = beratParsel % 1000

	// Menghitung biaya dasar berdasarkan per kg
	biayaPerKg = kg * 10000

	// Menghitung biaya tambahan berdasarkan sisa berat
	if kg > 10 {
		biayaSisa = 0 // Jika berat lebih dari 10 kg, sisa berat digratiskan
	} else if sisaBerat >= 500 {
		biayaSisa = sisaBerat * 5 // Jika sisa >= 500 gram, biaya tambahan Rp. 5 per gram
	} else {
		biayaSisa = sisaBerat * 15 // Jika sisa < 500 gram, biaya tambahan Rp. 15 per gram
	}

	// Menampilkan detail berat
	fmt.Printf("Detail berat: %d kg %d gr\n", kg, sisaBerat)

	// Menampilkan detail biaya
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaPerKg, biayaSisa)

	// Menampilkan total biaya
	fmt.Printf("Total biaya: Rp. %d\n", biayaPerKg+biayaSisa)
}
