package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		n      = 5
		bunga  string
		Pita   []string
		banyak int
	)
	//fmt.Print("N : ")
	//fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		fmt.Printf("Bunga %d: ", i)
		fmt.Scanln(&bunga)
		if bunga == "Selesai" || bunga == "selesai" {
			break
		}
		Pita = append(Pita, bunga)
		banyak++
	}
	temp := strings.Join(Pita, " - ")
	fmt.Println("Pita : " + temp + " - ")
	fmt.Print("Bunga : ", banyak)
}
