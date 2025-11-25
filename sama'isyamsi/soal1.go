package main
import "fmt"
func main() {
	var gram int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&gram)
	kg := gram/ 1000
	sisa := gram % 1000
	fmt.Printf("Detail berat: %d kg + &d gr/n", kg, sisa)
	biayaDasar := kg * 10000
	biayaTambahan := 0
	if kg >= 10 {
		biayaTambahan = 0
	} else {
		if sisa >= 500 {
			biayaTambahan = sisa * 5
		} else {
			biayaTambahan = sisa * 15
		}
	}
	total := biayaDasar + biayaTambahan
	fmt.Printf("Biaya dasar: Rp. %d + Rp.%d\n", biayaDasar, biayaTambahan)
	fmt.Printf("Total biaya: Rp. %d\n", total)

}