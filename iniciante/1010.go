//----> 1010. Cálculo Simples
package main
 
import (
    "fmt"
)
 
func main() {
    //Declaração de Variáveis
    var cod1, cod2 int64
    var num1, num2, valor1, valor2 float64
    
    //Entrada de dados
    fmt.Scanf("%d", &cod1)
    fmt.Scanf("%f", &num1)
    fmt.Scanf("%f", &valor1)
    fmt.Scanf("%d", &cod2)
    fmt.Scanf("%f", &num2)
    fmt.Scanf("%f", &valor2)
    
    //Cálculo do valor a pagar
    TOTAL := (num1 * valor1) + (num2 * valor2) 
    
    //Saída de dados
    fmt.Printf("VALOR A PAGAR: R$ %.2f\n", TOTAL) 
}

