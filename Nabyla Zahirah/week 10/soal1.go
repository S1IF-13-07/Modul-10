package main
import "fmt"
func main(){
	var gram int 
	fmt.Print("Masukkan berat parsel (gram): ")
	fmt.Scan(&gram)

    kg := gram / 1000
    sisa := gram % 1000
    biayaKg := kg * 10000

    var tambahan int
    if gram > 10000 {
        tambahan = 0
    } else if sisa >= 500 {
        tambahan = sisa * 5
    } else if sisa > 0 {
        tambahan = sisa * 15
    } else {
        tambahan = 0
    }

    total := biayaKg + tambahan
    fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisa)
    fmt.Printf("Detail biaya: Rp %d", biayaKg)
    if tambahan > 0 {
        fmt.Printf(" + Rp %d\n", tambahan)
    } else {
        fmt.Printf(" + Rp %d (sisa gratis)\n", tambahan)
    }
    fmt.Printf("Total biaya: Rp %d\n", total)
}
