package main

import "fmt"

func main(){
	var huruf rune
	fmt.Print("masukkan huruf :")
	fmt.Scanf("%c", &huruf)
	if huruf == 'a' || huruf == 'i' || huruf == 'u' || huruf == 'e' || huruf == 'o' || huruf == 'A' || huruf == 'I' || huruf == 'U' || huruf == 'E' || huruf == 'O' {
		fmt.Print("huruf vokal")
	} else if (huruf >= 'a' && huruf <= 'z') || (huruf >= 'A' && huruf <= 'Z') {
		fmt.Print("huruf konsonan")
	} else {
		fmt.Print("bukan huruf")
	}
}