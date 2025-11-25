package main
import "fmt"
func main() {
	
	var huruf rune

	fmt.Print("Masukkan sebuah huruf : ")
	fmt.Scanf("%c", &huruf)

	if huruf == 'A' || huruf == 'I' || huruf == 'U' || huruf == 'E' || huruf == 'O' ||
	 huruf == 'a' || huruf == 'i' || huruf == 'u' || huruf == 'e' || huruf == 'o' {
		fmt.Println("Vokal")
	} else if (huruf >= 'A' && huruf <= 'Z') || (huruf >= 'a' && huruf <= 'z') {
		fmt.Println("Konsonan")
	} else {
		fmt.Println("Bukan huruf")
	}
}