package main

import "fmt"

func main() {
	var berat_gram int
	fmt.Scan(&berat_gram)

	berat_kg := berat_gram / 1000
	sisa_gram := berat_gram % 1000

	biaya_kg := berat_kg * 10000

	var biaya_sisa int

	if berat_kg > 10 {
		biaya_sisa = 0

	} else {

		if sisa_gram >= 500 {
			biaya_sisa = sisa_gram * 5

		} else {
			biaya_sisa = sisa_gram * 15
		}
	}

	biaya_total := biaya_kg + biaya_sisa

	fmt.Println("Detail berat:", berat_kg, "kg +", sisa_gram, "gr")
	fmt.Println("Detail biaya: Rp.", biaya_kg, "+ Rp.", biaya_sisa)
	fmt.Println("Total biaya: Rp.", biaya_total)
}
