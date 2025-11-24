package main
import "fmt"
func main(){
 	var bilangan, d1, d2, d3, d4 int
 	var urutan string
 	
	fmt.Print("Bilangan: ")
 	fmt.Scan(&bilangan)
	d1 = bilangan / 1000
	d2 = (bilangan % 1000) / 100
	d3 = (bilangan % 100) / 10
	d4 = bilangan % 10
 
 	if d1 < d2 && d2 < d3 && d3 < d4 {
 		urutan = "terurut membesar"
 	} else if d1 > d2 && d2 > d3 && d3 > d4{
 		urutan = "terurut mengecil"
 	} else {
 		urutan = "tidak terurut"
 	}
 	fmt.Println("Digit pada bilangan", bilangan, urutan)
}