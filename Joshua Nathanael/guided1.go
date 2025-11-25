package main 

import "fmt"

func main() {
	var umur int
	var kk bool
	fmt.Print("Masukkan umur: ")
	fmt.Scan(&umur)
	fmt.Print("Apakah sudah memiliki kk? : ")
	fmt.Scan(&kk)
	if umur >= 17 && kk == true {
		fmt.Println("bisa membuat KTP")
	} else {
		fmt.Println("belum bisa membuat ktp")
	}
}