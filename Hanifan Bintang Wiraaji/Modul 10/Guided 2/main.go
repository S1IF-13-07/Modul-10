package main

import "fmt"

func main() {
	var c rune

	fmt.Scanf("%c", &c)

	if c == 'a' || c == 'i' || c == 'u' || c == 'e' || c == 'o' ||  
	c == 'A' || c == 'I' || c == 'U' || c == 'E' || c == 'O' {
		fmt.Println("vokal")
	} else if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') { 
		fmt.Println("konsonan")
	} else {
		fmt.Println("bukan huruf")
	}
}
