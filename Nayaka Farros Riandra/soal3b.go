package main

import "fmt"

func main() {
	var bilangan int
	var prima bool
	var jumlahFaktor int = 0

	fmt.Print("Bilangan : ")
	fmt.Scan(&bilangan)

	fmt.Print("Faktor : ")
	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Print(i, " ")
			jumlahFaktor++
		}
	}
	
	fmt.Println()

	if jumlahFaktor == 2 {
		prima = true
	} else {
		prima = false
	}

	fmt.Println("Prima:", prima)
}