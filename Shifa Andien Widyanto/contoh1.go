package main
import "fmt"
func main() {
	var umur int
	var KK bool

	fmt.Print("Masukan umur : ")
	fmt.Scan(&umur)
	fmt.Print("Memiliki KK : ")
	fmt.Scan(&KK)

	if umur >= 17 && KK == true {
		fmt.Println("bisa membuat KTP")
	} else {
		fmt.Println("belum bisa membuat KTP")
	}
}