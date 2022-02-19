//----> 1014. Consumo
package main
 
import (
    "fmt"
)
 
func main() {
    //Declaração de variáveis
    var X, Y float64 
    
    //Entrada de dados
    fmt.Scanf("%f", &X)
    fmt.Scanf("%f", &Y)
    
    //Cálculo do consumo médio
    media := X / Y 
    
    //Saída de dados 
    fmt.Printf("%.3f km/l\n", media)
}
