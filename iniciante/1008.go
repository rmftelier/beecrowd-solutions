//----> 1008. Salário
package main
 
import (
    "fmt"
)
 
func main() {
    //Declaração de Variáveis
    var numero int64
    var qtdHoras, valor float64
    
    //Entrada de dados
    fmt.Scanf("%d", &numero)
    fmt.Scanf("%f", &qtdHoras)
    fmt.Scanf("%f", &valor)
    
    //Cálculo do salário
    SALARY := qtdHoras * valor
    
    //Saída de dados
    fmt.Printf("NUMBER = %d\n", numero)
    fmt.Printf("SALARY = U$ %.2f\n", SALARY) 
}