package main

import "fmt"

func main() {
	var bilangan, d1, d2, d3, d4 int
	fmt.Print("Bilangan: ")
	fmt.Scan(&bilangan)
	d4 = bilangan % 10
	d3 = (bilangan % 100) / 10
	d2 = (bilangan % 1000) / 100
	d1 = bilangan / 1000
	if d1 < d2 && d2 < d3 && d3 < d4 {
		fmt.Print("Terurut Membesar")
	} else if d1 > d2 && d2 > d3 && d3 > d4 {
		fmt.Print("Terurut Mengecil")
	} else {
		fmt.Print("Tidak Terurut")
	}
}
