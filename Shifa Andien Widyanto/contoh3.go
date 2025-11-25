package main  
import "fmt"  
func main() {  
    var bilangan, a1, a2, a3, a4 int 
    var teks string 
    fmt.Print("Bilangan: ") 
    fmt.Scan(&bilangan) 
    a4 = bilangan % 10 
    a3 = (bilangan / 10) % 10 
    a2 = (bilangan / 100) % 10 
    a1 = bilangan / 1000 
    if a1 < a2 && a2 < a3 && a3 < a4 {  
        teks = "terurut membesar" 
    }else if a1 > a2 && a2 > a3 && a3 > a4{ 
        teks = "terurut mengecil" 
    }else{  
        teks = "tidak terurut" 
    } 
    fmt.Println("Digit pada bilangan", bilangan, teks) 
} 