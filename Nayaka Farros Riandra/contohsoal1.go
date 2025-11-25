package main

import "fmt"

func main() {
	var umur int
	var b bool
	fmt.Print("masukkan umur :")
	fmt.Scan(&umur)
	fmt.Print("punya KK? (true/false) :")
	fmt.Scan(&b)
	if umur >= 17 && b == true {
		fmt.Print("bisa membuat ktp")
	} else {
		fmt.Print("belum bisa membuat ktp")
	}
}