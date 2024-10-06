package main

import (
	"fmt"
	"strings"
)

func main() {
	var bunga string
	var pita []string
	var jumlahBunga int

	for {
		// Meminta input nama bunga
		fmt.Printf("Bunga %d: ", jumlahBunga+1)
		fmt.Scan(&bunga)  // Membaca input nama bunga

		// Cek jika input adalah "SELESAI" (tidak case-sensitive)
		if strings.ToUpper(bunga) == "SELESAI" {
			break
		}

		// Menambahkan bunga ke dalam pita
		pita = append(pita, bunga)
		jumlahBunga++
	}

	// Tampilkan pita dan jumlah bunga setelah selesai
	if jumlahBunga > 0 {
		result := strings.Join(pita, " – ")
		fmt.Println("Pita:", result)
		fmt.Printf("Bunga: %d\n", jumlahBunga)
	} else {
		fmt.Println("Pita: ")
		fmt.Println("Bunga: 0")
	}
}