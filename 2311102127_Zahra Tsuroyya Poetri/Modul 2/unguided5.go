package main

import (
	"fmt"
)

func main() {
	var b int

	// Meminta input bilangan bulat b > 1
	fmt.Print("Bilangan: ")
	fmt.Scanln(&b)

	// Memastikan input valid
	if b <= 1 {
		fmt.Println("Bilangan harus lebih dari 1.")
		return
	}

	// Menampilkan faktor-faktor bilangan
	fmt.Print("Faktor: ")
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i," ")
		}
	}
	fmt.Println() // Untuk memberikan newline di akhir output
}
