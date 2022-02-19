//----> 1007. Diferença
package main
 
import (
    "fmt"
)
 
func main() {
	  //Declaração de variáveis
    var A, B, C, D int64
    
		//Entrada de dados
    fmt.Scanf("%d", &A)
    fmt.Scanf("%d", &B)
    fmt.Scanf("%d", &C)
    fmt.Scanf("%d", &D)

    
    //Cálculo da diferença
    DIFERENCA := (A * B - C * D)
    
    //Saída de dados
    fmt.Printf("DIFERENCA = %d\n", DIFERENCA)
}
