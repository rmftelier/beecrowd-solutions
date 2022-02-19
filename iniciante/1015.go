//----> 1015. Distância Entre Dois Pontos 
package main
 
import (
    "fmt"
    "math"
)
 
func main() {    
    //Declaração de variáveis
    var x1, y1, x2, y2 float64

    //Entrada de dados
    fmt.Scanf("%f", &x1)
    fmt.Scanf("%f", &y1)
    fmt.Scanf("%f", &x2)
    fmt.Scanf("%f", &y2)

    //Cálculo da distância
    distancia := math.Sqrt(math.Pow((x2 - x1), 2.0) + math.Pow((y2 - y1), 2.0))
    
    //Saída de dados 
    fmt.Printf("%.4f\n", distancia)
}