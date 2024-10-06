/*Liya Khoirunnisa - 2311102124*/

package main

import (
	"bufio"   // Membaca input dari pengguna
	"fmt"     // Menampilkan teks ke layar
	"os"      // Mendapatkan input dari terminal
	"strings" // Mengolah teks
)

func main() {
	// Membaca input
	reader := bufio.NewReader(os.Stdin)

	// Deklarasi pita untuk menyimpan daftar bunga
	var pita []string
	nomorBunga := 1

	// Perulangan inputan bunga
	for {
		fmt.Printf("Bunga %d: ", nomorBunga)
		bunga, _ := reader.ReadString('\n')
		bunga = strings.TrimSpace(bunga)

		// Perulangan akan berhenti jika input "selesai" 
		if strings.ToLower(bunga) == "selesai" {
			break
		}

		// Jika tidak maka disimpan di pita
		pita = append(pita, bunga)
		nomorBunga++
	}

	// Menampilkan hasil
	if len(pita) > 0 {
		fmt.Printf("Pita: %s -\n", strings.Join(pita, " - "))
		fmt.Printf("Bunga: %d\n", len(pita))
	} else {
		fmt.Println("Pita: ")
		fmt.Println("Bunga: 0")
	}
}
