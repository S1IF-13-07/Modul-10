package main
import "fmt"
func main() {
    var beratGram int
    fmt.Print("Berat parsel (gram): ")
    fmt.Scan(&beratGram)

    kg := beratGram / 1000
    sisa := beratGram % 1000

    biayaKg := kg * 10000
    biayaSisa := 0

    if kg > 10 {
        biayaSisa = 0
    } else {
        if sisa >= 500 {
            biayaSisa = sisa * 5
        } else {
            biayaSisa = sisa * 15
        }
    }

    total := biayaKg + biayaSisa
    fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisa)
    fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaSisa)
    fmt.Printf("Total biaya: Rp. %d\n", total)
}
