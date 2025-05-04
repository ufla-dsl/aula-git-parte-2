package main

import (
	"fmt"
	"os"
	"strconv"

	"tempconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run ./app <temperatura>")
		os.Exit(1)
	}

	input, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Erro: valor inválido")
		os.Exit(1)
	}

	c := tempconv.Celsius(input)
	f := tempconv.Fahrenheit(input)

	fmt.Printf("%s = %s\n", c, tempconv.CToF(c))
	fmt.Printf("%s = %s\n", f, tempconv.FToC(f))
}
