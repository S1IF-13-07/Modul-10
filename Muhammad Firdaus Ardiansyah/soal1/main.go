package main

import "fmt"

func main(){
	var gram, kilogram, parcel, totalbiaya, biayaSisa int

	fmt.Print("berat parcel : ")
	fmt.Scan(&parcel)

	kilogram = parcel / 1000
	gram = parcel % 1000
	detailbiaya := kilogram * 10000
	
	fmt.Println("detail berat :", kilogram, "kg", "+", gram, "gr")
	if gram >= 500{
		biayaSisa := gram * 5
		totalbiaya = detailbiaya + biayaSisa
		fmt.Println("Detail biaya : Rp. ", detailbiaya, "+", "Rp. ", biayaSisa)
	} else if gram < 500{
		biayaSisa = parcel * 15
		totalbiaya = detailbiaya + biayaSisa
		fmt.Println("Detail biayaSisa : Rp. ", detailbiaya, "+", "Rp. ", biayaSisa)
	} else if gram >= 10000{
		biayaSisa = detailbiaya
		fmt.Println("Detail biayaSisa :Rp. ", detailbiaya, "+", "Rp. ", biayaSisa)
	}
	fmt.Println("total biaya : Rp. ", totalbiaya)
}