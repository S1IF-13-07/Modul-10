package main

import (
	"fmt"
)

func main() {
	var char rune
	fmt.Scanf("%c", &char)
	if char == 'a' || char == 'A' || char == 'i' || char == 'I' || char == 'u' || char == 'U' || char == 'e' || char == 'E' || char == 'o' || char == 'O' {
		fmt.Print("vokal")
	} else if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
		fmt.Print("konsonan")
	} else {
		fmt.Print("Bukan Huruf")
	}
}
