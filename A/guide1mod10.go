package main
import "fmt"
func main(){
	var usia int
	var kk bool
	fmt.Scan(&usia,&kk)
	if usia >=17 && kk == true {
		fmt.Println ("bisa membuat ktp")
	}else{
		fmt.Println ("tidak bisa membuat ktp")
	}
}