package main

import "fmt"

func main() {
	var bilangan, d1, d2, d3, d4 int
	fmt.Print("Masukkan sebuah bilangan 4 digit: ")
	fmt.Scan(&bilangan)

	d1 = bilangan / 1000
	d2 = (bilangan / 100) % 10
	d3 = (bilangan / 10) % 10
	d4= bilangan % 10

	if d1 < d2 && d2 < d3 && d3 < d4 {
		fmt.Println("Bilangan tersebut naik")
	} else if d1 > d2 && d2 > d3 && d3 > d4 {
		fmt.Println("Bilangan tersebut turun")
	} else {
		fmt.Println("Bilangan tersebut tidak naik dan tidak turun")
	}
}