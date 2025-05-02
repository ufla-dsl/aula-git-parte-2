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
	temp, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Printf("Erro ao converter %q: %v\n", os.Args[1], err)
		os.Exit(1)
	}

	// Converte e exibe a temperatura nas diferentes escalas
	c := tempconv.Celsius(temp)
	f := tempconv.Fahrenheit(temp)

	// Para a saída em Celsius, formata com duas casas decimais
	cToF := tempconv.CToF(c)
	fToC := tempconv.FToC(f)

	fmt.Printf("%.0f°C = %.0f°F\n", c, cToF)
	fmt.Printf("%.0f°F = %.2f°C\n", f, fToC)
}
