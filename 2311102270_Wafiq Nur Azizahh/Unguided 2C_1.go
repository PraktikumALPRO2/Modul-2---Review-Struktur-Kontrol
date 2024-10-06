package main

import (
	"fmt"
)

func main() {
	var (
		total                 int
		sisa, diskon, hargakg int
	)

	fmt.Print("Berat persel (gram) : ")
	fmt.Scanln(&total)
	kg := total / 1000
	grm := total % 1000
	fmt.Printf("Detail berat : %d kg + %d gram", kg, grm)

	if total > 10000 {
		hargakg = kg * 10000
		if grm >= 500 {
			sisa = grm * 5
		} else if grm < 500 {
			sisa = grm * 15
		}
		fmt.Println("\nDetail biaya : Rp.", hargakg, " + Rp.", sisa)
		if sisa <= 1000 {
			diskon = 0
		}
		totalbiaya := hargakg + diskon
		fmt.Println("Total biaya : Rp.", totalbiaya)
	} else {
		hargakg = kg * 10000
		if grm >= 500 {
			sisa = grm * 5
		} else if grm < 500 {
			sisa = grm * 15
		}
		fmt.Println("\nDetail biaya : Rp.", hargakg, " + Rp.", sisa)
		totalbiaya := hargakg + sisa
		fmt.Println("Total biaya : Rp.", totalbiaya)
	}

}
