package main

import (
	"fmt"
)

func main() {
	var berat1, berat2 float64

	for {
		// Meminta input berat masing-masing kantong
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&berat1, &berat2)

		// Jika salah satu kantong mencapai atau melebihi 9 kg, hentikan program
		if berat1 >= 9 || berat2 >= 9 {
			fmt.Println("Proses selesai.")
			break
		}
	}
}
