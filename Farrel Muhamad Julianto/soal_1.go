package main

import "fmt"

func main() {
	var beratGram int
	var kg, sisaGram int
	var biayaPerKg, biayaSisa, totalBiaya int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&beratGram)
	kg = beratGram / 1000
	sisaGram = beratGram % 1000
	biayaPerKg = kg * 10000
	if sisaGram >= 500 {
		biayaSisa = sisaGram * 5
	} else {
		biayaSisa = sisaGram * 15
	}
	if kg > 10 {
		biayaSisa = 0
	}
	totalBiaya = biayaPerKg + biayaSisa
	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaPerKg,
		biayaSisa)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
