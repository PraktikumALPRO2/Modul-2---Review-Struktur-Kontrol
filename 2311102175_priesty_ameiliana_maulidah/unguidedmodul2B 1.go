package main

import (
	"fmt"
)

func main() {
	var colors [5][4]string
	successSequence := [4]string{"merah", "kuning", "hijau", "ungu"}
	success := true

	for i := 0; i < 5; i++ {
		fmt.Printf("Masukkan warna untuk percobaan ke-%d (4 warna): ", i+1)
		for j := 0; j < 4; j++ {
			fmt.Scan(&colors[i][j])
		}
		if colors[i] != successSequence {
			success = false
		}
	}

	fmt.Println(success)
}
