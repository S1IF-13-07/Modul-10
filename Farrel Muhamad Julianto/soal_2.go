package main

import "fmt"

func main() {
	var nilai float64
	fmt.Print("Masukkan nilai akhir: ")
	fmt.Scan(&nilai)

	var grade string

	switch {
	case nilai >= 81:
		grade = "A"
	case nilai > 72.5:
		grade = "AB"
	case nilai > 65:
		grade = "B"
	case nilai > 57.5:
		grade = "BC"
	case nilai > 50:
		grade = "C"
	case nilai > 40:
		grade = "D"
	default:
		grade = "E"
	}

	fmt.Printf("Hasil penilaian: %s\n", grade)
}
