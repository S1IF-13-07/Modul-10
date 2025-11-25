package main

import "fmt"

func main() {
	var umur int
	var kk bool
	fmt.Print("Masukkan umur: ")
	fmt.Scan(&umur)
	fmt.Print("punya KK?: ")
	fmt.Scan(&kk)

	if umur >= 17 && kk {
		fmt.Print("bisa membuat KTP")
	} else {
		fmt.Print("belum bisa membuat KTP")
	}
}