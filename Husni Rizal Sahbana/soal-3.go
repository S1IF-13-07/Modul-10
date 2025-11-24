package main

import "fmt"

func main() {
	var x int
	var hitung_faktor int
	var prima bool

	fmt.Print("Masukan Angka : ")
	fmt.Scan(&x)

	fmt.Print("Factorial: ")

	for i := 1; i <= x; i++ {
		if x%i == 0 {
			fmt.Print(i, " ")
			hitung_faktor = hitung_faktor + 1
		}
	}

	fmt.Println()

	if hitung_faktor == 2 {
		prima = true
	} else {
		prima = false
	}

	fmt.Println("Prima:", prima)
}
