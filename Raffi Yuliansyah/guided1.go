package main

import "fmt"

func main() {
	var umur int
	var kk bool
	fmt.Scan(&umur, &kk)
	if umur >= 17 && kk == true {
		fmt.Print("Bisa membuat ktp")
		} else {
		fmt.Print("Belum bisa membuat ktp")
	}
}
