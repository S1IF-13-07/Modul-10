package main

import "fmt"

func main() {
	var umur int
	var kartu_keluarga bool
	fmt.Scan(&umur, &kartu_keluarga)

	if umur > 17 && kartu_keluarga {
		fmt.Println("Bisa membuat KTP")
	} else {
		fmt.Println("Maaf belum bisa membuat KTP")
	}
}
