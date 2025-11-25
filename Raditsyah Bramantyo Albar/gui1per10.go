package main

import "fmt"

func main() {
    var umur int
    var ktp bool

    fmt.Print("Masukkan umur: ")
    fmt.Scanln(&umur)
    fmt.Print("Apakah memiliki kartu keluarga (true/false): ")
    fmt.Scanln(&ktp)
    
    if umur >= 17 && ktp {
        fmt.Println("bisa membuat KTP")
    } else {
        fmt.Println("belum bisa membuat KTP")
    }
}
