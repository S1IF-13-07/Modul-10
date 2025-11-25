package main

import "fmt"

func main() {
	var x int
	var y bool

	fmt.Print("Masukkan usia: ")
	fmt.Scan(&x)

	fmt.Print("Apakah mempunyai KK? (True/False): ")
	fmt.Scan(&y)

	if x >= 17 && y == true {
		fmt.Println("Bisa membuat KTP")
	} else {
		fmt.Println("Belum bisa membuat KTP")
	}

	//cara 2
	// if x >= 17 && y == true {
    //     fmt.Println("Bisa membuat KTP")
    // } else if x >= 17 && y == false {
    //     fmt.Println("Belum bisa membuat KTP")
    // } else {
    //     fmt.Println("Belum bisa membuat KTP")
    // }
}
