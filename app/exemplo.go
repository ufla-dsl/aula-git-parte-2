package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ufla-gcc259/aula-git-parte-2/tempconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Uso: go run ./app <temperatura>")
		os.Exit(1)
	}

	// Converte o argumento para float64
	valor, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Printf("Valor inválido: %v\n", os.Args[1])
		os.Exit(1)
	}

	// Converte o valor para Celsius e Fahrenheit
	c := tempconv.Celsius(valor)
	f := tempconv.Fahrenheit(valor)

	// Exibe os resultados das conversões
	fmt.Printf("%v = %v\n", c, tempconv.CToF(c))
	fmt.Printf("%v = %v\n", f, tempconv.FToC(f))
}
