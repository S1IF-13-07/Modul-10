package main
import "fmt"
func main() {
	var usia int
	var punyakk bool
	fmt.Scan(&usia)
	fmt.Scan(&punyakk)
	if usia >= 17 && punyakk {
		fmt.Println("bisa membuat ktp")
	} else {
		fmt.Println("belum bisa membuat ktp")
	}

}