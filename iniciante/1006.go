//----> 1006. Média 2
package main
 
import (
    "fmt"
)
 
func main() {
    //Declaração de variáveis
    var A, B, C float64 
    
		//Entrada de dados
    fmt.Scanf("%f", &A)
    fmt.Scanf("%f", &B)
    fmt.Scanf("%f", &C)
    
    //Cálculo da média
    MEDIA := ((2 * A) + (3 * B) + (5 * C)) / 10
    
    //Saída de dados
    fmt.Printf("MEDIA = %.1f\n", MEDIA)
}