//----> 1017. Gasto de Combustível 
package main
 
import (
    "fmt"
)
 
func main() {
    //Declaração de variáveis
    var horas, velocidade float64 //Tempo gasto - horas | Velocidade média - km/h

    //Entrada de dados 
    fmt.Scanf("%f", &horas)
    fmt.Scanf("%f", &velocidade) 
    
    //Cálculo da qntd de litros de combustível gastos em uma viagem
    litrosNec := (horas * velocidade)/ 12.0
    
    //Saída de dados 
    fmt.Printf("%.3f\n", litrosNec) 
}
