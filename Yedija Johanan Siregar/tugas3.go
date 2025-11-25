 package main
 import "fmt"
 
  func main() {
   var b int
   fmt.Print("Masukkan sebuah bilangan bulat: ")
   fmt.Scan(&b)
   fmt.Print("Faktor: ")
 
  for i := 1; i <= b; i++ {
  if b%i == 0 {
   fmt.Print(i, " ")
   }
 }
   fmt.Println()
     count := 0
     for i := 1; i <= b; i++ {
     if b%i == 0 {
     count++
   }
 }
   prima := count == 2
   fmt.Println("Apakah bilangan prima?", prima)
 }
