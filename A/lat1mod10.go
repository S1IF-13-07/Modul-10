package main

import "fmt"

func main() {
	var berat int
	fmt.Print("Berat parcel (gram): ")
	fmt.Scan(&berat)

	kg := berat / 1000
	sisa := berat % 1000

	fmt.Println("Detail berat :", kg, "kg +", sisa, "gr")

	biaya := kg * 10000

	if kg >= 10 {
		fmt.Println("Detail biaya: Rp", kg*10000)
		fmt.Println("Total biaya :", biaya)
		return
	}

	if sisa < 500 {
		fmt.Println("Detail biaya: Rp", kg*10000, "+ Rp", sisa*15)
		biaya += sisa * 15
	} else {
		fmt.Println("Detail biaya: Rp", kg*10000, "+ Rp", sisa*5)
		biaya += sisa * 5
	}

	fmt.Println("Total biaya :", biaya)
}