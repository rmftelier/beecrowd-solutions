//----> 1011. Esfera
package main
 
import (
    "fmt"
    "math"
)

func main() {
	  //Declaração de variáveis
    var raio float64

		//Entrada de dados
    fmt.Scanf("%f", &raio)

		//Cálculo do volume
    volume := (4.0 / 3) * 3.14159 * math.Pow(raio, 3.0)
    
		//Saída de dados 
		fmt.Printf("VOLUME = %.3f\n", volume)
}
