// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import (
	"fmt"
)

func hitungBiayaPengiriman(beratGram int) {
	// Menghitung berat dalam kg dan sisa gram
	kg := beratGram / 1000
	gram := beratGram % 1000

	// Biaya dasar (Rp 10.000 per kg)
	biayaDasar := kg * 10000

	// Menghitung biaya tambahan untuk sisa gram
	var biayaTambahan int
	if gram <= 500 {
		biayaTambahan = gram * 5
	} else {
		biayaTambahan = gram * 15
	}

	// Total biaya
	totalBiaya := biayaDasar + biayaTambahan

	// Membatasi biaya maksimal untuk berat lebih dari 10 kg
	if kg >= 10 {
		totalBiaya = 100000
	}

	// Menampilkan hasil
	fmt.Printf("Berat paket (gram): %d\n", beratGram)
	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, gram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaDasar, biayaTambahan)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}

func main() {
	// Contoh penggunaan fungsi
	hitungBiayaPengiriman(8500)
	hitungBiayaPengiriman(9250)
	hitungBiayaPengiriman(11750)
}
