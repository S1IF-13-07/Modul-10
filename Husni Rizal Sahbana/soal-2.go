package main

import "fmt"

func main() {
	var nilai float64
	var predikat string

	fmt.Print("Masukkan Nilai Akhir : ")
	fmt.Scan(&nilai)

	if nilai > 80 {
		predikat = "A"

	} else if nilai > 72.5 {
		predikat = "AB"

	} else if nilai > 65 {
		predikat = "B"

	} else if nilai > 57.5 {
		predikat = "BC"

	} else if nilai > 50 {
		predikat = "C"

	} else if nilai > 40 {
		predikat = "D"

	} else {
		predikat = "E"
	}

	fmt.Println("Nilai akhir nya dalam bentuk predikat adalah : ", predikat)
}
