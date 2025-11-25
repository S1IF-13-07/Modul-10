package main
import "fmt"

func main() {
	var n, i, faktor int
	var prima bool
	fmt.Print("masukkan bilangan: ")
	fmt.Scan(&n)
	faktor = 0

	fmt.Print("faktor: ")
	for i = 1; i <= n; i++ {
		if n % i == 0 {
			fmt.Print(" ", i)
			faktor++
		} 
	}

	fmt.Println()

	fmt.Print("Prima: ")
	if faktor == 2 {
		prima = true
	}
	fmt.Print(prima)
}