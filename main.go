package main

import (
	"fmt"

	"github.com/ufla-gcc259/aula-git-parte-2/tempconv"
)

func main() {
	fmt.Printf("Que frio! %v\n", tempconv.AbsoluteZeroC)
	fmt.Printf("Fervendo! %v\n", tempconv.CToF(tempconv.BoilingC))

	// Caso de Uso: conversões CToK e FToK
	fmt.Println("\nConversões para Kelvin:")
	fmt.Println("0°C =", tempconv.CToK(0))
	fmt.Println("32°F =", tempconv.FToK(32))

	// Caso de Uso: conversões KToC e KToF
	fmt.Println("\nConversões de Kelvin:")
	fmt.Println("273.15K =", tempconv.KToC(273.15), "=", tempconv.KToF(273.15))
}
