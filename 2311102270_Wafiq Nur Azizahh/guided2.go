package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	correct0rder := []string{"merah", "kuning", "hijau", "ungu"}

	reader := bufio.NewReader(os.Stdin)
	success := true

	for i := 1; i <= 5; i++ {
		fmt.Print("Percobaan ", i, " : ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		colors := strings.Split(input, " ")

		for j := 0; j < 4; j++ {
			if colors[j] != correct0rder[j] {
				success = false
				break
			}
		}
		if  !success {
			break
		}
	}

	if success {
		fmt.Println("Berhasil : true")
	}else {
		fmt.Println("Berhasil : false")
	}
}
