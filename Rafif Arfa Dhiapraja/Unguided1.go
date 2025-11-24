package main

import "fmt"

func main() {
	var berat,hargaSisa int
	fmt.Scan(&berat)

	kg := berat / 1000
	sisa := berat % 1000

	hargaKg := kg * 10000

	if kg > 10 {
		hargaSisa = 0
	} else if sisa >= 500 {
		hargaSisa = sisa * 5
	} else if sisa < 500 {
		hargaSisa = sisa * 15
	}

	totalHarga := hargaKg + hargaSisa

	fmt.Println("Detail Berat:", kg, "kg", "+", sisa, "gr")
	fmt.Println("Detail Biaya: Rp.", hargaKg, "+ Rp.", hargaSisa)
	fmt.Println("Total Biaya: Rp.", totalHarga)
}