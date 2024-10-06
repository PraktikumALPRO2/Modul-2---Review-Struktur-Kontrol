// lutfiana Isnaeni Lathifah
// 2311102165

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var pita string
	var count int

	// Loop hingga user mengetikkan "SELESAI"
	for {
		fmt.Printf("Bunga %d: ", count+1)
		bunga, _ := reader.ReadString('\n')
		bunga = strings.TrimSpace(bunga)

		// Jika input adalah "SELESAI", keluar dari loop
		if strings.ToUpper(bunga) == "SELESAI" {
			break
		}

		if count == 0 {
			pita = bunga
		} else {
			pita = pita + " - " + bunga
		}
		count++
	}

	// Tampilkan hasil pita dan jumlah bunga
	fmt.Println("Pita:", pita)
	fmt.Println("Bunga:", count)
}
