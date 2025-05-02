package main

import (
	"fmt"
	"github.com/ufla-gcc259/aula-git-parte-2/tempconv"
)

func main() {
	var opcao int
	var valor float64

	fmt.Println("Escolha uma opção:")
	fmt.Println("1 - Celsius para Fahrenheit")
	fmt.Println("2 - Fahrenheit para Celsius")
	fmt.Print("> ")
	fmt.Scanln(&opcao)

	fmt.Print("\nDigite o valor: ")
	fmt.Scanln(&valor)

	switch opcao {
	case 1:
		c := tempconv.Celsius(valor)
		fmt.Printf("%.2f°C = %.2f°F\n", c, tempconv.CToF(c))
	case 2:
		f := tempconv.Fahrenheit(valor)
		fmt.Printf("%.2f°F = %.2f°C\n", f, tempconv.FToC(f))
	default:
		fmt.Println("Opção inválida.")
	}
}
