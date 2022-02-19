//----> 1038. Lanche
package main
 
import (
    "fmt"
)
 
func main() {
    //Declaração de variáveis 
    var codigo, qtd int64
    var total float64
    
    //Entrada de dados 
    fmt.Scanf("%d", &codigo)
    fmt.Scanf("%d", &qtd)

    //Conversão do tipo int para float 
	  var qtnd float64 = float64(qtd)
  
    //Usando Switch
    switch codigo {
     case 1:
        total = 4.00 * qtnd
     case 2:
        total = 4.50 * qtnd
     case 3:
        total = 5.00 * qtnd
     case 4:
        total = 2.00 * qtnd
     case 5:
        total = 1.50 * qtnd
     default:
        total = 0.00
     }   

		 //Saída de dados
     fmt.Printf("Total: R$ %.2f\n", total)
}
