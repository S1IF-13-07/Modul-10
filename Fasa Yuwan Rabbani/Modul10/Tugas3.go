package main

import "fmt"

func main() {
	var b int
	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	fmt.Print("Faktor: ")
	factorCount := 0
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Printf("%d ", i)
			factorCount++
		}
	}
	fmt.Println()

	prima := factorCount == 2
	fmt.Printf("Prima: %v\n", prima)
}
