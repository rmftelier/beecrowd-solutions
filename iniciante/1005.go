//----> 1005. Média 1
package main
 
import (
    "fmt"
)
 
func main() {
    //Declaração de variáveis 
    var A, B float64 
    
		//Entrada de dados 
    fmt.Scanf("%f", &A)
    fmt.Scanf("%f", &B)
    
    //Cálculo da média
    MEDIA := ((3.5 * A) + (7.5 * B))/11
    
    //Saída de dados 
    fmt.Printf("MEDIA = %.5f\n", MEDIA)
}