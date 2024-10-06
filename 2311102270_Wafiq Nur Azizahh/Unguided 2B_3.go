package main

import "fmt"

func main() {
	var (
		a, b float32
	)
	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong : ")
		fmt.Scanln(&a, &b)
		// if a == 9 || b == 9 {
		// 	break
		// } else if a-b >= 9 {
		// 	break
		// }
		if a < b {
			a, b = b, a
		}

		if a-b <= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng : false")
		} else if a+b >= 150 || a < 0 || b < 0 {
			break
		} else if a-b >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng : true")
		}
	}
	fmt.Print("Proses Selesai.")
}
