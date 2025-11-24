package main

import "fmt"

func main() {
	var n int
	var kk bool
	fmt.Scan(&n)
	fmt.Scan(&kk)
	if n >= 17 && kk {
		fmt.Print("Bisa membuat ktp")
	} else {
		fmt.Print("Belum bisa membuat")
	}
}
