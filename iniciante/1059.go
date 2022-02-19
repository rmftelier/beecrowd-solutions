//----> 1059. Números pares
package main
 
import (
    "fmt"
)
 
func main() {
     //Processamento e saída de dados 
     for contador := 1; contador <= 100; contador ++{
         if contador % 2 == 0{
             fmt.Printf("%d\n", contador)
         }
     }
}
