//----> 1067. Números ímpares
package main
 
import (
    "fmt"
)
 
func main() {
    //Declaração de variáveis 
    var X int 
    
    //Entrada de dados 
    fmt.Scanf("%d", &X)
    
    //Processamento e saída de dados 
   for contador := 1; contador <= X; contador++ {
       if contador % 2 != 0 {
              fmt.Printf("%d\n", contador)
       }
   }
}
