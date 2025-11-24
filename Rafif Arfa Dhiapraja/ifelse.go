package main

import "fmt"

func main() {
	var usia int
	var KK bool
	fmt.Print("Masukkan usia Anda: ")
	fmt.Scan(&usia)
	fmt.Print("Apakah Anda memiliki Kartu Keluarga (true/false): ")
	fmt.Scan(&KK)

	if usia >= 17 && KK == true {
		fmt.Println("Bisa Membuat KTP")
	} else {
		fmt.Println("Tidak Bisa Membuat KTP")
	}

}