package main

import (
	"fmt"
)

func main() {
	var b int
	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	if b <= 1 {
		fmt.Println("Bilangan harus lebih dari 1.")
		return
	}

	factors := []int{}
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			factors = append(factors, i)
		}
	}

	fmt.Print("Faktor: ")
	for _, factor := range factors {
		fmt.Print(factor, " ")
	}
	fmt.Println()

	isPrime := len(factors) == 2
	fmt.Println("Prima:", isPrime)
}

