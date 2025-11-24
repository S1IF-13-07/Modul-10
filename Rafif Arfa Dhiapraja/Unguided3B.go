package main

import "fmt"

func main() {
	var b int
	var prima bool
	var JumlahFaktor int = 0

	fmt.Print("bilangan: ")
	fmt.Scan(&b)

	fmt.Print("Faktor: ")
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
			JumlahFaktor++
			} 
		}

		fmt.Println()

		if JumlahFaktor == 2 {
			prima = true
		}

		fmt.Print("prima: ", prima)

	


}