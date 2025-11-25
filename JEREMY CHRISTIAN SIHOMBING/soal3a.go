package main
import "fmt"
func main(){

	var bilangan int

	fmt.Print("bilangan : ")
	fmt.Scan(&bilangan)
	fmt.Print("Faktor : ")
	
	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Print(i, " ")
		}
	}
}