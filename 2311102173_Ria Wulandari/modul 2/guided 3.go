package main

import "fmt"

func main() {
	// Deklarasi variabel untuk menyimpan angka
	var a, b, c, d, e int
	var hasil int

	// Minta input dari pengguna
	fmt.Scanln(&a, &b, &c, &d, &e)

	// Hitung hasil penjumlahan
	hasil = a + b + c + d + e

	// Tampilkan hasil penjumlahan
	fmt.Printf("Hasil Penjumlahan %d, %d, %d, %d, %d adalah = %d\n", a, b, c, d, e, hasil)
}
