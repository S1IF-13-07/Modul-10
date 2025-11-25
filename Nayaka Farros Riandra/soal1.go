package main

import "fmt"

func main() {
	var beratParsel, kg, gram, biayaKg, biayaSisa, totalBiaya int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&beratParsel)

	kg = beratParsel / 1000
	gram = beratParsel % 1000

	biayaKg = kg * 10000

	if kg >= 10 {
		biayaSisa = 0
	} else if gram >= 500 {
		biayaSisa = gram * 5
	} else {
		biayaSisa = gram * 15
	}

	totalBiaya = biayaKg + biayaSisa


	fmt.Println("Detail berat:", kg, "kg +", gram, "gr")
	fmt.Println("Detail biaya: Rp.", biayaKg, "+ Rp.", biayaSisa)
	fmt.Println("Total biaya: Rp.", totalBiaya)
}