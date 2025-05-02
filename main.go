package main

import (
	"fmt"
	"os"
	"strconv"

	"raynnertaniguch1 /tempconv/tempconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Uso: go run ./app <temperatura>")
		return
	}

	valor, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Erro: argumento inválido. Use um número.")
		return
	}

	c := tempconv.Celsius(valor)
	f := tempconv.Fahrenheit(valor)

	fmt.Printf("%s = %s\n", c, tempconv.CToF(c))
	fmt.Printf("%s = %s\n", f, tempconv.FToC(f))
}
