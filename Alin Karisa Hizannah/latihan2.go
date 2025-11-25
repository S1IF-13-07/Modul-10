package main

import "fmt"

func main() {
	var input rune

	fmt.Print("Masukkan satu karakter: ")
	fmt.Scan(&input)

	if input == 'a' || input == 'i' || input == 'u' || input == 'e' || input == 'o' ||
		input == 'A' || input == 'I' || input == 'U' || input == 'E' || input == 'O' {

		fmt.Println("Vokal")

	} else if (input >= 'a' && input <= 'z') || (input >= 'A' && input <= 'Z') {

		fmt.Println("Konsonan")

	} else {

		fmt.Println("Bukan huruf")
	}
}
