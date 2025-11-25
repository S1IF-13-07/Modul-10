package main
import "fmt"
func main() {
	var a int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&a)

	beratKg := a / 1000
	sisaBerat := a % 1000

	biayaDasar := beratKg * 10000
	var biayaTambahan int
	if sisaBerat >= 500 {
		biayaTambahan = sisaBerat * 5
	} else if sisaBerat > 0 {
		biayaTambahan = sisaBerat * 15
	}
	var totalBiaya int
	if a > 10000 {
		totalBiaya = biayaDasar
	} else {
		totalBiaya = biayaDasar + biayaTambahan
	}
	fmt.Println()
	fmt.Println("Detail berat: ", beratKg, "kg +", sisaBerat, "gr")
	fmt.Println("Detail biaya: Rp.", biayaDasar, "+ Rp.", biayaTambahan)
	fmt.Println("Total biaya: Rp.", totalBiaya)
}
