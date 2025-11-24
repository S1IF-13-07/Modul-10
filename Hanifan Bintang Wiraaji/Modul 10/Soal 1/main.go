package main

import "fmt"

func main() {
	var b int

	fmt.Print("Berat parsel(gram): ")
	fmt.Scan(&b)

	kgberat := b / 1000
	gberat := b % 1000
	fmt.Printf("Detail berat: %v kg + %v gr\n", kgberat, gberat)

	kgbiaya := kgberat * 10000
	gbiaya := gberat
	if gbiaya >= 500 {
		gbiaya = gbiaya * 5
	} else {
		gbiaya = gbiaya * 15
	}
	fmt.Printf("Detail biaya: Rp. %v + Rp. %v\n", kgbiaya, gbiaya)
	
	var totalbiaya int
	if b > 10000 {
		totalbiaya = kgbiaya
	} else {
		totalbiaya = kgbiaya + gbiaya
	}
	fmt.Printf("Total biaya: Rp. %v", totalbiaya)
}
