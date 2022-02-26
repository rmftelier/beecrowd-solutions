//----> 1048. Aumento de Salário
package main
 
import (
    "fmt"
)
 
func main() {
    
    //Declaração de variáveis
    var salario, reajuste, sal_nov float64

    //Entrada de dados
    fmt.Scanf("%f", &salario)
    
    //Cálculo do Reajuste Salarial | Processamento e Saída de Dados
    if salario >= 0 && salario <= 400.00 {
        
        reajuste = salario * 0.15 
        sal_nov = salario + reajuste 
        
        fmt.Printf("Novo salario: %.2f\n", sal_nov)
        fmt.Printf("Reajuste ganho: %.2f\n", reajuste)
        fmt.Print("Em percentual: 15 %\n")
    } else if salario >= 400.01 && salario <= 800.00{
        
        reajuste = salario * 0.12 
        sal_nov = salario + reajuste 
      
        fmt.Printf("Novo salario: %.2f\n", sal_nov)
        fmt.Printf("Reajuste ganho: %.2f\n", reajuste)
        fmt.Print("Em percentual: 12 %\n")
    } else if salario >- 800.01 && salario <= 1200.00{
        
        reajuste = salario * 0.10
        sal_nov = salario + reajuste

        fmt.Printf("Novo salario: %.2f\n", sal_nov)
        fmt.Printf("Reajuste ganho: %.2f\n", reajuste)
        fmt.Print("Em percentual: 10 %\n")
                   
    } else if salario >= 1200.01 && salario <= 2000.00{
        
        reajuste = salario * 0.07
        sal_nov = salario + reajuste
        
        fmt.Printf("Novo salario: %.2f\n", sal_nov)
        fmt.Printf("Reajuste ganho: %.2f\n", reajuste)
        fmt.Print("Em percentual: 7 %\n")
    } else {
        
        reajuste = salario * 0.04
        sal_nov = salario + reajuste 
        fmt.Printf("Novo salario: %.2f\n", sal_nov)
        fmt.Printf("Reajuste ganho: %.2f\n", reajuste)
        fmt.Print("Em percentual: 4 %\n")
    }
}