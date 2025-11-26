package main

import "fmt"

func main() {
	var bilangan int

	fmt.Print("Bilangan: ")
	fmt.Scan(&bilangan)

	fmt.Print("Faktor: ")
	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()

	jumlahFaktor := 0
	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			jumlahFaktor++
		}
	}

	prima := jumlahFaktor == 2
	fmt.Println("Prima:", prima)
}
