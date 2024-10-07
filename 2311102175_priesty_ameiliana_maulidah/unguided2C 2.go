package main

import "fmt"

func main () {
	var nam float32
	var nmk string

	// meminta input nilai
	fmt.Print("masukkan niali: ")
	fmt.Scan(&nam)

	// logika penentuan nilai huruf berdasarkan nilai numerik
	if nam> 80 {
		nmk = "A"
	} else if nam > 75.5{
		nmk = "B"
	} else if nam > 65 {
		nmk = "c"
	}else if nam > 50 {
		nmk = "D"
	}else if nam > 40 {
		nmk = "E"
	}else{
		nmk = "f"
	}

	// menampilkan hasil 
	fmt.Printf("nilai indeks untuk nilai %.2f adalah %s\n", nam, nmk)
}