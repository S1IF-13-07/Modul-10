package main
import "fmt"
func main(){
	var usia int
	var a bool
	fmt.Scan(&usia)
	fmt.Scan(&a)
	if usia >= 17 && a == true{
		fmt.Print("bisa membuat ktp")
	} else {
		fmt.Print("belum bisa memmbuat ktp")
	}
}