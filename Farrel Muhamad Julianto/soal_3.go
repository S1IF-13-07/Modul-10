package main

import "fmt"

func main() {
    var b int
    fmt.Print("Bilangan: ")
    fmt.Scan(&b)

    fmt.Print("Faktor: ")
    factors := []int{}
    for i := 1; i <= b; i++ {
        if b%i == 0 {
            factors = append(factors, i)
            fmt.Print(i, " ")
        }
    }
    fmt.Println()

    isPrime := false
    if b > 1 && len(factors) == 2 {
        isPrime = true
    }

    fmt.Println("Prima:", isPrime)
}
