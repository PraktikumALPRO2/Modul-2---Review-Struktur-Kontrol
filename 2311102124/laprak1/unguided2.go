/*Liya Khoirunnisa - 2311102124*/

package main

import "fmt" // Berfungsi untuk menampilkan teks ke layar

func main() {
	var kantong1, kantong2 float64

	for {
		// Input berat belanjaan dari kedua kantong
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kantong1, &kantong2)

		// Cek jika salah satu kantong beratnya negatif atau total berat > 150
		if kantong1 < 0 || kantong2 < 0 || (kantong1+kantong2) > 150 {
			fmt.Println("Proses selesai")
			break
		}

		// Hitung apakah sepeda motor oleng atau tidak
		if kantong1 > kantong2 {
			if kantong1-kantong2 >= 9 {
				fmt.Println("Sepeda motor pak Andi akan oleng: true")
			} else {
				fmt.Println("Sepeda motor pak Andi akan oleng: false")
			}
		} else {
			if kantong2-kantong1 >= 9 {
				fmt.Println("Sepeda motor pak Andi akan oleng: true")
			} else {
				fmt.Println("Sepeda motor pak Andi akan oleng: false")
			}
		}
	}
}
