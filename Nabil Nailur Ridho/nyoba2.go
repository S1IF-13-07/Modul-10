package main
import "fmt"
func main(){
	var huruf rune
	fmt.Print("Masukkan Huruf: ")
	fmt.Scanf("%c", &huruf)

    if huruf == 'A' || huruf == 'I' || huruf == 'U' || huruf == 'E' || huruf == 'O' || huruf == 'a' || huruf == 'i' || huruf == 'u' || huruf == 'e' || huruf == 'o' {
        fmt.Print("vokal")
    } else if (huruf >= 'a' && huruf <= 'z') || (huruf >= 'A' && huruf <= 'Z') {
        fmt.Print("konsonan")
    } else {
		fmt.Print("bukan huruf")
	}
}