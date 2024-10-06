package main

import (
	"fmt"
)

func main() {
	// Input berat dalam gram
	var beratGram int
	fmt.Print("Masukkan berat parcel (gram): ")
	fmt.Scanln(&beratGram)

	// Konversi ke kg dan gram sisa
	beratKg := beratGram / 1000
	sisaGram := beratGram % 1000

	// Hitung biaya kirim dasar per kg
	biayaPerKg := 10000
	totalBiaya := beratKg * biayaPerKg

	// Variabel tambahan biaya
	var biayaTambahan int

	// Hitung biaya tambahan
	if beratKg >= 10 {
		// Jika berat lebih dari atau sama dengan 10 kg, tidak ada biaya tambahan
		biayaTambahan = 0
	} else if sisaGram <= 500 {
		// Jika sisa gram <= 500, biaya tambahan dihitung Rp 5 per gram
		biayaTambahan = 2500
	} else {
		// Jika sisa gram > 500, biaya tambahan dihitung Rp 15 per gram
		biayaTambahan = 3750
	}

	// Tambahkan biaya tambahan ke total biaya
	totalBiaya += biayaTambahan

	// Tampilkan hasil sesuai format contoh
	fmt.Println("\n=============================")
	fmt.Printf("Berat parcel (gram): %d\n", beratGram)
	fmt.Printf("Detail berat: %d kg + %d gr\n", beratKg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", beratKg*biayaPerKg, biayaTambahan)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
	fmt.Println("=============================\n")
}
