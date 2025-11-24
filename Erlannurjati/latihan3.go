package main
import "fmt"
func main() {
    var b int
    fmt.Print("Bilangan: ")
    fmt.Scan(&b)

    fmt.Print("Faktor: ")

    jumlahFaktor := 0

    for f := 1; f <= b; f++ {
        if b%f == 0 {
            fmt.Print(f, " ")
            jumlahFaktor++
        }
    }
    prima := false
    if jumlahFaktor == 2 {
        prima = true
    }

    fmt.Println()
    fmt.Println("Prima:", prima)
}
