//----> 1004. Produto Simples
package main
 
import (
    "fmt"
)
 
func main() {  
    //Declaração de variáveis 
    var A, B int 
    
    //Entrada de dados
    fmt.Scanf("%d", &A)
    fmt.Scanf("%d", &B) 
    
    //Cálculo do produto 
    X := A * B
    
    //Saída de dados 
    fmt.Printf("PROD = %d\n", X)
}