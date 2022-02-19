//----> 1035. Teste de Seleção 1
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
    
    //Processamento e saída de dados 
    if B > C && D > A && (C + D) > (A + B) && C > 0 && D > 0 && A % 2 == 0 {
        fmt.Printf("Valores aceitos\n")
    } else {
        fmt.Printf("Valores nao aceitos\n") 
    }
}
