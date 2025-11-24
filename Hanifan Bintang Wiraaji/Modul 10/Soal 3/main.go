package main

import "fmt"

func main() {
	var b, prima int
	fmt.Print("Bilangan: ")
	fmt.Scan(&b)
	if b > 1 {
		fmt.Print("Faktor: ")
		for i := 1; i <= b; i++ {
			faktor := b % i
			if faktor == 0 {
				fmt.Print(i, " ")
				prima += i
			}
		}
		fmt.Println("")
	}
	if b > 0 {
		bilPrima := false
		prima -= 1
		if prima == b {
			bilPrima = true
		}
		fmt.Printf("Prima: %v", bilPrima)
	}
}
