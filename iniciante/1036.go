//----> 1036. Fórmula de Bhaskara
package main
 
import (
    "fmt"
		"math"
)
 
func main() {
    //Declaração de variáveis
		var A, B, C, x1, x2 float64

		//Entrada de dados
    fmt.Scanf("%f", &A)
    fmt.Scanf("%f", &B)
    fmt.Scanf("%f", &C)

		//Cálculo do Delta
		delta := math.Pow(B, 2) - (4 * A * C)

		//Processamento e Saída de dados 
		if delta < 0 || A == 0 {
         
			fmt.Print("Impossivel calcular\n") 
	 }else{
			x1 = ((- B) + math.Sqrt(delta))/(2 * A) 
			x2 = ((- B) - math.Sqrt(delta))/(2 * A)    
			
			fmt.Printf("R1 = %.5f\n", x1)
			fmt.Printf("R2 = %.5f\n", x2)
	 }
}