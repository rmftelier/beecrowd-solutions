//----> 1009. Salário com Bônus
package main
 
import (
    "fmt"
)
 
func main() {
    //Declaração de variáveis
    var nome string 
    var salFixo, montVendas float64
       
    //Entrada de dados 
    fmt.Scanf("%s", &nome)
    fmt.Scanf("%f", &salFixo)
    fmt.Scanf("%f", &montVendas) 
        
    //Cálculo do salário
    TOTAL := salFixo + (montVendas * 0.15)
    
    //Saída de dados
    fmt.Printf("TOTAL = R$ %.2f\n", TOTAL) 
}