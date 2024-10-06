package main

import "fmt"

func main() {
	var bil1, bil2, i int
	var prima1, prima2 bool

	fmt.Print("Bilangan 1: ")
	fmt.Scanln(&bil1)

	fmt.Print("Faktor: ")
	prima1 = true
	for i = 1; i <= bil1; i++ {
		if bil1%i == 0 {
			fmt.Print(i, " ")
			if i != 1 && i != bil1 {
				prima1 = false
			}
		}
	}
	fmt.Println()

	fmt.Print("Bilangan 2: ")
	fmt.Scanln(&bil2)

	fmt.Print("Faktor: ")
	prima2 = true
	for i = 1; i <= bil2; i++ {
		if bil2%i == 0 {
			fmt.Print(i, " ")
			if i != 1 && i != bil2 {
				prima2 = false
			}
		}
	}
	fmt.Println()

	if bil1 == 1 {
		prima1 = false
	}
	if bil2 == 1 {
		prima2 = false
	}

	fmt.Println("Prima 1:", prima1)
	fmt.Println("Prima 2:", prima2)
}
