// package main

// import "fmt"

// func main() {
// 	var (
// 		b      int
// 		faktor []int
// 	)
// 	fmt.Print("Bilangan: ")
// 	fmt.Scanln(&b)
// 	if b > 0 {
// 		for i := 1; i <= b; i++ {
// 			if b%i == 0 {
// 				faktor = append(faktor, i)
// 			}
// 		}
// 		fmt.Print("Faktor: ")
// 		for j := 0; j < len(faktor); j++ {
// 			fmt.Print(faktor[j], " ")
// 		}
// 		if len(faktor) == 2 {
// 			fmt.Print("\nPrima: true")
// 		} else {
// 			fmt.Print("\nPrima: false")
// 		}

// 	}
// }

package main

import "fmt"

func main() {
	var nam float32
	var nmk string
	fmt.Print("Masukkan Nilai : ")
	fmt.Scan(&nam)

	if nam > 80 {
		nmk = "A"
	} else if nam > 72.5 {
		nmk = "B"
	} else if nam > 65 {
		nmk = "C"
	} else if nam > 50 {
		nmk = "D"
	} else if nam > 40 {
		nmk = "E"
	} else {
		nmk = "F"
	}
	fmt.Printf("Nilai Indeks untuk nilai %.2f adalah %s\n", nam, nmk)
}
