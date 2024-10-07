package main

import (
	"fmt"
)

func main() {
	var berat1, berat2 float64

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		_, err := fmt.Scan(&berat1, &berat2)
		if err != nil {
			fmt.Println("Input tidak valid.")
			continue
		}

		if berat1 < 0 || berat2 < 0 || (berat1+berat2) > 150 {
			break
		}

		selisih := berat1 - berat2
		oleng := selisih >= 9 || selisih <= -9
		fmt.Printf("Sepeda motor pak Andi akan oleng: %t\n", oleng)
	}

	fmt.Println("Proses selesai.")
}
