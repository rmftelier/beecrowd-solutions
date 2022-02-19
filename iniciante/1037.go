//----> 1037. Intervalo
package main
 
import (
    "fmt"
)
 
func main() {  
    //Declaração de variáveis 
    var num float64 
    
    //Entrada de dados 
    fmt.Scanf("%f", &num)
    
    //Processamento e saída de dados 
    if num >= 0 && num <= 25 {
       fmt.Printf("Intervalo [0,25]\n")
    } else if num > 25 && num <= 50 {
        fmt.Printf("Intervalo (25,50]\n")
    } else if  num > 50 && num <= 75 {
        fmt.Printf("Intervalo (50,75]\n")        
    } else if num > 75 && num <= 100 {
        fmt.Printf("Intervalo (75,100]\n")
    } else {
        fmt.Printf("Fora de intervalo\n")
    }

}
