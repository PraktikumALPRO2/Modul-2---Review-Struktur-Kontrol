package main

import "fmt"

func hitungFaktorDanPrima(bil int) (bool, []int) {
	var faktor []int
	prima := true

	for i := 1; i <= bil; i++ {
		if bil%i == 0 {
			faktor = append(faktor, i)
			if i != 1 && i != bil {
				prima = false
			}
		}
	}
	return prima, faktor
}

func main() {
	var bil1, bil2 int

	fmt.Print("Bilangan 1: ")
	fmt.Scanln(&bil1)

	prima1, faktor1 := hitungFaktorDanPrima(bil1)
	fmt.Print("Faktor: ")
	for _, f := range faktor1 {
		fmt.Print(f, " ")
	}
	fmt.Println()
	fmt.Println("Prima 1:", prima1)

	fmt.Print("Bilangan 2: ")
	fmt.Scanln(&bil2)

	prima2, faktor2 := hitungFaktorDanPrima(bil2)
	fmt.Print("Faktor: ")
	for _, f := range faktor2 {
		fmt.Print(f, " ")
	}
	fmt.Println()

	if bil1 == 1 {
		prima1 = false
	}
	if bil2 == 1 {
		prima2 = false
	}

	fmt.Println("Prima 2:", prima2)
}
