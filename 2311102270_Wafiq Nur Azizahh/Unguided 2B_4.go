package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		K, f float64
	)

	fmt.Print("Nilai : ")
	fmt.Scanln(&K)

	f = math.Pow(4*K+2, 2) / ((4*K + 1) * (4*K + 3))
	fmt.Printf("Nilai f(k) = %.10f", f)

}
