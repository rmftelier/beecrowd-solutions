//----> 1013. O Maior
package main
 
import (
    "fmt"
)
 
func main() {
	  //Declaração de variáveis
    var valor1, valor2, valor3 int
    
		//Entrada de dados 
    fmt.Scanf("%d", &valor1)
    fmt.Scanf("%d", &valor2)
    fmt.Scanf("%d", &valor3)
    
		//Processamento e saída de dados
	  if valor1 > valor2 && valor1 > valor3 {
		   fmt.Printf("%d eh o maior\n", valor1)
	  } else if valor2 > valor3 {
		   fmt.Printf("%d eh o maior\n", valor2)
	  } else {
		   fmt.Printf("%d eh o maior\n", valor3)
	  }
}

