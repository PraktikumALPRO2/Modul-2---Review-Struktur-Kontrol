package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	fmt.Print("Masukkan jumlah bunga: ")
	fmt.Scan(&N)

	var pita string
	for i := 1; i <= N; i++ {
		var bunga string
		fmt.Printf("Bunga %d: ", i)
		fmt.Scan(&bunga)

		if bunga == "SELESAI" {
			break
		}

		if pita == "" {
			pita = bunga
		} else {
			pita += " " + bunga
		}
	}

	bungaList := strings.Split(pita, " ")
	fmt.Printf("Pita: %s\n", pita)
	fmt.Printf("Jumlah bunga: %d\n", len(bungaList))
}
