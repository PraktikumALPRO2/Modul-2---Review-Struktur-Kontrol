package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	var bunga string
	var pita []string

	// Meminta input jumlah bunga N
	fmt.Print("Masukkan jumlah bunga (N): ")
	fmt.Scan(&N)

	// Jika N lebih dari 0, proses input bunga
	if N > 0 {
		for i := 1; i <= N; i++ {
			fmt.Printf("Bunga %d: ", i)
			fmt.Scan(&bunga)  // Membaca input nama bunga
			pita = append(pita, bunga)
		}
		// Menggabungkan nama-nama bunga dengan separator " – "
		result := strings.Join(pita, " – ")
		fmt.Println("Pita:", result)
	} else {
		fmt.Println("Pita: ")
	}
}
