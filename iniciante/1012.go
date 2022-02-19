//----> 1012. Area
package main
 
import (
    "fmt"
)
 
func main() {
    //Declaração de variáveis
    var A, B, C float64

    //Entrada de dados 
    fmt.Scanf("%f", &A)
    fmt.Scanf("%f", &B)
    fmt.Scanf("%f", &C)

    //Área do triângulo retângulo
    TRIANGULO := (A * C)/ 2.0
    
    //Área do círculo de raio C 
    CIRCULO := 3.14159 * (C * C)
    
    //Área do trapézio
    TRAPEZIO := ((A + B) * C) / 2.0
    
    //Área do quadrado 
    QUADRADO := B * B
    
    //Área do retângulo 
    RETANGULO := A * B 

    //Saída de dados
    fmt.Printf("TRIANGULO: %.3f\n", TRIANGULO)
    fmt.Printf("CIRCULO: %.3f\n", CIRCULO)
    fmt.Printf("TRAPEZIO: %.3f\n", TRAPEZIO)
    fmt.Printf("QUADRADO: %.3f\n", QUADRADO)
    fmt.Printf("RETANGULO: %.3f\n", RETANGULO)
}
