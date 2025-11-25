package main

import "fmt"

func main() {
	var ch rune

	fmt.Print("Masukkan satu karakter: ")
	fmt.Scanf("%c", &ch)

	if ch == 'a' || ch == 'i' || ch == 'u' || ch == 'e' || ch == 'o' ||
		ch == 'A' || ch == 'I' || ch == 'U' || ch == 'E' || ch == 'O' {

		fmt.Println("Vokal")

	} else if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {

		fmt.Println("Konsonan")
		
	} else {

		fmt.Println("Bukan huruf")
	}
}
	
