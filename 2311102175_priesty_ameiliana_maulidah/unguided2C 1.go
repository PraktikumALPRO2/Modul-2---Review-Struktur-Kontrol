package main

import (
	"fmt"
)

func main() {
	var beratParcel int
	fmt.Print("Masukkan berat parsel dalam gram: ")
	fmt.Scan(&beratParcel)

	totalBeratKg := beratParcel / 1000
	sisaBeratGram := beratParcel % 1000

	biayaPengiriman := totalBeratKg * 10000

	if totalBeratKg > 10 {
		sisaBeratGram = 0
	} else if sisaBeratGram >= 500 {
		biayaPengiriman += sisaBeratGram * 5
	} else {
		biayaPengiriman += sisaBeratGram * 15
	}

	fmt.Printf("Total biaya pengiriman: Rp. %d\n", biayaPengiriman)
}

