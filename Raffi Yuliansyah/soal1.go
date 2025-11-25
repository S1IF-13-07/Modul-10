package main

import "fmt"

func main() {
	var berat int
	fmt.Print("Berat parsel (gram) : ")
	fmt.Scan(&berat)

	fmt.Println("Detai berat : ", berat/1000, " kg", " + ", berat%1000, " gram")
	if berat >= 10000 {
		if berat%1000 >= 500 {
			fmt.Println("Detail Biaya : ", "Rp. ", (berat/1000)*10000, " + ", "Rp. ", (berat%1000)*5)
			fmt.Println("Total Biaya : ", "Rp. ", ((berat / 1000) * 10000))
		} else {
			fmt.Println("Detail Biaya : ", "Rp. ", (berat/1000)*10000, " + ", "Rp. ", (berat%1000)*15)
			fmt.Println("Total Biaya : ", "Rp. ", ((berat / 1000) * 10000))
		}
	} else if berat%1000 >= 500 {
		fmt.Println("Detail Biaya : ", "Rp. ", (berat/1000)*10000, " + ", "Rp. ", (berat%1000)*5)
		fmt.Println("Total Biaya : ", "Rp. ", ((berat/1000)*10000)+((berat%1000)*5))
	} else {
		fmt.Println("Detail Biaya : ", "Rp. ", (berat/1000)*10000, " + ", "Rp. ", (berat%1000)*15)
		fmt.Println("Total Biaya : ", "Rp. ", ((berat/1000)*10000)+((berat%1000)*15))
	}
}
