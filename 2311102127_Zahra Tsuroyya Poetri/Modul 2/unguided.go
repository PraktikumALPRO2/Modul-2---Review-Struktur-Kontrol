package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Membaca input untuk N (jumlah input bunga)
	var N int
	fmt.Print("N: ")
	fmt.Scanln(&N)

	// Jika N = 0, langsung tampilkan pita kosong
	if N == 0 {
		fmt.Println("Pita: ")
		return
	}

	// Membaca inputan bunga dari pengguna
	reader := bufio.NewReader(os.Stdin)
	pita := " "

	// Loop untuk menerima input bunga sebanyak N kali
	for i := 1; i <= N; i++ {
		fmt.Printf("Bunga %d: ", i)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Menggabungkan input bunga ke dalam pita
		pita += input + " - "
	}

	// Menampilkan hasil akhir pita
	fmt.Printf("Pita: %s\n", pita)
}


