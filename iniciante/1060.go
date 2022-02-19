//----> 1060. Números Positivos 
package main
 
import (
    "fmt"
)
 
func main() {

    //Declaração de variáveis 
    var num float64
    pos := 0 //Quantidade de números positivos 
    
    //Entrada e processamento de dados
    for contador := 1; contador <= 6; contador++ {
        fmt.Scanf("%f", &num)
        
        if num > 0 {
            pos++
        }
    }
    
    //Saída de dados 
    fmt.Printf("%d valores positivos\n", pos)    
}
