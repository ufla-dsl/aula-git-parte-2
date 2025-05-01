*tempconv*
=====
Biblioteca escrita em Go para conversões simples de temperatura. Extraída do livro **A Linguagem de Programação Go**, de Alan A. A. Donovan e Brian Kernighan. 

Como usar?
----
Para rodar o programa localmente e entender o que ele faz, siga os passos abaixo:

1. Clone este repositório, entre no seu diretório e rode o programa

git clone https://github.com/ufla-dsl/aula-git-parte-2.git

cd aula-git-parte-2

go run main.go

```go
package main

import (
	"fmt"

	"github.com/ufla-gcc259/aula-git-parte-2/tempconv"
)

func main() {
	fmt.Printf("Que frio! %v\n", tempconv.AbsoluteZeroC) // Que frio! -273.15°C
	fmt.Printf("Fervendo! %v\n", tempconv.CToF(tempconv.BoilingC)) // Fervendo! 212°F
}
```

Aplicação de exemplo
----
O repositório inclui uma aplicação de exemplo (`app.go`) que demonstra a conversão de temperaturas entre as escalas Celsius e Fahrenheit.

Para executar a aplicação, forneça um valor de temperatura como parâmetro:

```
go run ./app 100
```

Saída esperada:
```
100°C = 212°F
100°F = 37.78°C
```

O código da aplicação:

```go
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
```

Outras linguagens?
----
Versões da biblioteca *tempconv* para outras linguagens:

> *Todo*


Licença
-----

> *Todo*


Como contribuir?
----
Escolha uma *issue* dentre as [disponíveis](https://github.com/ufla-gcc259/aula-git-parte-2/issues), avise à comunidade que você está trabalhando nela e envie um *Pull Request*, quando terminar.
