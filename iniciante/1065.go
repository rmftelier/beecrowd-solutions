//----> 1065. Pares entre Cinco Números
package main
 
import (
    "fmt"
)
 
func main() {
    //Declaração de variáveis 
    var num int64
    pares := 0 //Quantidade de números pares 
    
    //Leitura de dados 
    for contador := 1; contador <= 5; contador++ {
        fmt.Scanf("%d", &num)

        if num % 2 == 0 {
            pares++
        }
    }
    
    //Saída de dados 
    fmt.Printf("%d valores pares\n", pares)   
}