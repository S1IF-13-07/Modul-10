package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	angka1 := n / 1000
	angka2 := n % 1000 / 100
	angka3 := n % 1000 % 100 / 10
	angka4 := n % 10

	if angka1 <= angka2 && angka2 <= angka3 && angka3 <= angka4 {
		fmt.Printf("Digit pada bilangan %v terurut membesar", n)
	} else if angka1 >= angka2 && angka2 >= angka3 && angka3 >= angka4 {
		fmt.Printf("Digit pada bilangan %v terurut mengecil", n)
	}else {
		fmt.Printf("Digit pada bilangan %v tidak terurut", n)
	}
}
