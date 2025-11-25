package main

import "fmt"

func main() {
	var bilangan int
	var prima bool
	prima = true
	fmt.Print("Bilangan : ")
	fmt.Scan(&bilangan)
	fmt.Print("Faktor : ")
	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Print(i, " ")
		}
	}

	if bilangan == 0 || bilangan == 1 {
		prima = false
	}

	for i := 2; i <= bilangan-1; i++ {
		if bilangan%i == 0 {
			prima = false
			break
		}
	}

	if prima {
		fmt.Print("Prima : ", prima)
	} else {
		fmt.Print("Prima : ", prima)
	}
}
