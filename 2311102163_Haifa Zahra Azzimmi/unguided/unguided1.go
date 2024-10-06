package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var jumlahBunga int
	pita := ""
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("N: ")
	fmt.Scanln(&jumlahBunga)

	for i := 1; i <= jumlahBunga; i++ {
		fmt.Printf("Bunga %d: ", i)
		scanner.Scan()
		bunga := scanner.Text()

		if pita == "" {
			pita = bunga
		} else {
			pita += " - " + bunga
		}
	}

	fmt.Printf("Pita: %s\n", pita)
}
