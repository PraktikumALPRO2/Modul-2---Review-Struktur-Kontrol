package main

import (
	"fmt"
)

func main() {
	var beratKantong1, beratKantong2 float64
	var totalBerat float64
	selisih := 0.0
	const batasTotalBerat = 150.0

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&beratKantong1, &beratKantong2)

		totalBerat = beratKantong1 + beratKantong2

		if beratKantong1 < 0 || beratKantong2 < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		if totalBerat > batasTotalBerat {
			fmt.Println("Total berat lebih dari 150 kg. Proses selesai.")
			break
		}

		if beratKantong1 > beratKantong2 {
			selisih = beratKantong1 - beratKantong2
		} else {
			selisih = beratKantong2 - beratKantong1
		}

		if selisih >= 9 {
			fmt.Println("Sepeda motor Pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor Pak Andi akan oleng: false")
		}
	}
}
