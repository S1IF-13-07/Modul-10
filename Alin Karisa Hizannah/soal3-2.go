package main

import "fmt"

func main() {
    var b int
    var jumlahFaktor int = 0

    fmt.Print("Bilangan: ")
    fmt.Scan(&b)

    fmt.Print("Faktor: ")
    for i := 1; i <= b; i++ {
        if b%i == 0 {
            fmt.Print(i, " ")
            jumlahFaktor++
        }
    }
    fmt.Println()

    fmt.Print("Prima: ")
    if jumlahFaktor == 2 {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
}
