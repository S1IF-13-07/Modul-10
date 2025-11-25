package main

import "fmt"

func main() {
	var parsel, kilogram, gram, harga, biaya, total int
	fmt.Print("Berat parsel(gram): ")
	fmt.Scan(&parsel)

	kilogram = parsel / 1000
	gram = parsel % 1000
	harga = kilogram * 10000

	if parsel > 10000 {
		biaya = gram * 5
		total = harga
	} else if gram >= 500 {
		biaya = gram * 5
		total = harga + biaya
	} else if gram < 500 {
		biaya = gram * 15
		total = harga + biaya
	}

	fmt.Printf("detail berat: %d kg + %d gr\n", kilogram, gram)
	fmt.Printf("detail biaya: Rp. %d + Rp. %d\n", harga, biaya)
	fmt.Printf("total biaya: Rp. %d", total)
}
