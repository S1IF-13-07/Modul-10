package main
import "fmt"
func main() {
	var a int
	
	fmt.Print("Bilangan : ")
	fmt.Scan(&a)
	fmt.Print("Faktor : ")

	Faktor := 0
	for i := 1; i <= a; i++ {
		if a%i == 0 {
			fmt.Print(i, " ")
			Faktor++
		}
	}
	prima := (Faktor == 2)

	fmt.Println()
	fmt.Println("prima : ", prima)
}