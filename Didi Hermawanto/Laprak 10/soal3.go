package main

import "fmt"

func main() {
	var b int

	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	fmt.Print("Faktor: ")

	faktor := []int{}
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			faktor = append(faktor, i)
			fmt.Print(i, " ")
		}
	}

	prima := len(faktor) == 2

	fmt.Println()
	fmt.Println("Prima:", prima)
}
