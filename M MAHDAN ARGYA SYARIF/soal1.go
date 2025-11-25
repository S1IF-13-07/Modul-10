package main
import "fmt"

func main() {
	var berat, kg, g, harga, biaya, total int
	fmt.Print("Berat parsel(gram): ")
	fmt.Scan(&berat)

	kg = berat / 1000
	g = berat % 1000
	harga = kg * 10000

	if berat > 10000 {
		biaya = g * 5
		total = harga
	} else if g >= 500 {
		biaya = g * 5
		total = harga + biaya
	} else if g < 500 {
		biaya = g * 15
		total = harga + biaya
	} 

	fmt.Printf("detail berat: %d kg + %d gr\n", kg, g)
	fmt.Printf("detail biaya: Rp. %d + Rp. %d\n", harga, biaya)
	fmt.Printf("total biaya: Rp. %d", total)
}