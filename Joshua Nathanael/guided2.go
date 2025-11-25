package main

import "fmt"

func main() {
	var a rune

	fmt.Print("Masukkan suatu huruf: ")
	fmt.Scanf("%c", &a)

	if a == 'A' || a == 'I' || a == 'U' || a == 'E' || a == 'O' || a == 'a' || a == 'i' || a == 'u' || a == 'e' || a == 'o' {
		fmt.Println("Vokal")
	} else if (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z') {
		fmt.Println("konsonan")
	} else {
		fmt.Println("bukan huruf")
	}
}