//----> 1002. Área do Círculo
package main
 
import (
    "fmt"
)
 
func main() {     
	  //Declaração de variáveis 
    var raio float64
    
		//Entrada de dados
    fmt.Scanf("%f", &raio)  
		  
		//Fórmula do cálculo da área do círculo
    area := 3.14159 * (raio * raio)
    
		//Saída de dados 
    fmt.Printf("A=%.4f\n", area)
}