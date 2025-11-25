package main

import "fmt"

func main() {
	var angka, d1, d2, d3, d4 int
	var urutan string

	fmt.Print("Masukkan angka nya mas: ")
	fmt.Scan(&angka)

	d1 = angka / 1000
	d2 = (angka % 1000) / 100
	d3 = (angka %100) / 10
	d4 = angka % 10

	if d1 < d2 && d2 < d3 && d3 < d4 {
		urutan = "terurut membesar"
	} else if d1 > d2 && d2 > d3 && d3 > d4 {
		urutan = "terurut mengecil"
	} else {
		urutan = "tidak terurut"
	}
	fmt.Println("digit pada bilangan" , angka," ", urutan)
}