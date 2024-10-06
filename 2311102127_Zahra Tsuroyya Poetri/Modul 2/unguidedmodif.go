package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Membaca inputan bunga dari pengguna
	reader := bufio.NewReader(os.Stdin)
	pita := ""
	count := 0

	// Loop menerima inputan bunga, sampai "SELESAI" dimasukkan untuk berhenti
	for {
		fmt.Printf("Bunga %d: ", count+1)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Jika pengguna mengetikkan "SELESAI", hentikan input
		if strings.ToUpper(input) == "SELESAI" {
			break
		}

		// Menggabungkan input bunga ke dalam pita
		pita += input + " - "
		count++
	}

	// Menampilkan hasil akhir pita dan jumlah bunga
	fmt.Printf("Pita: %s\n", pita)
	fmt.Printf("Bunga: %d\n", count)
}
