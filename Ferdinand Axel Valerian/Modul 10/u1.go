package main

import "fmt"

func main() {
    var beratGram int
    
    fmt.Print("Berat parsel (gram): ")
    fmt.Scan(&beratGram)
    
    kg := beratGram / 1000
    sisaGram := beratGram % 1000
    
    biayaDasar := kg * 10000
    
    var biayaTambahan int
    
    if kg > 10 {

        biayaTambahan = 0
    } else if sisaGram >= 500 {
        biayaTambahan = sisaGram * 5
    } else {
        biayaTambahan = sisaGram * 15
    }
    
    fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisaGram)
    
    if kg > 10 {
        fmt.Printf("Detail biaya: Rp. %d + Rp. 0\n", biayaDasar)
    } else {
        fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaDasar, biayaTambahan)
    }
    
    totalBiaya := biayaDasar + biayaTambahan
    fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}