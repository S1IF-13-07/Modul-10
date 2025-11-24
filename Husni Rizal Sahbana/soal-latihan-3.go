package main

import "fmt"

func main() {
	var angka, d1, d2, d3, d4 int

	fmt.Scan(&angka)

	d4 = angka % 10           
	d3 = (angka % 100) / 10   
	d2 = (angka % 1000) / 100 
	d1 = angka / 1000         

	if d1 < d2 && d2 < d3 && d3 < d4 {
		fmt.Println("terurut membesar")
	} else if d1 > d2 && d2 > d3 && d3 > d4 {
		fmt.Println("terurut mengecil")
	} else {
		fmt.Println("tidak berurutan")
	}
}
