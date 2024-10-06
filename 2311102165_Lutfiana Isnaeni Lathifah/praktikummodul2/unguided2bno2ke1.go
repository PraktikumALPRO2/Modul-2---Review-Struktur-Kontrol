//Lutfiana Isnaeni Lathifah
//2311102165

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Meminta input N (jumlah bunga)
	fmt.Print("N: ")
	inputN, _ := reader.ReadString('\n')
	N, _ := strconv.Atoi(strings.TrimSpace(inputN))

	var pita string

	// Loop untuk memasukkan bunga
	for i := 1; i <= N; i++ {
		fmt.Printf("Bunga %d: ", i)
		bunga, _ := reader.ReadString('\n')
		bunga = strings.TrimSpace(bunga)
		if i == 1 {
			pita = bunga
		} else {
			pita = pita + " - " + bunga
		}
	}

	// Tampilkan hasil pita
	fmt.Println("Pita:", pita)
}
