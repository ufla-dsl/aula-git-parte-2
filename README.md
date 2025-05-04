# tempconv

Biblioteca escrita em Go para conversões simples de temperatura. Extraída do livro **A Linguagem de Programação Go**, de Alan A. A. Donovan e Brian Kernighan.

## Como usar?

Para rodar o programa localmente e entender o que ele faz, siga os passos abaixo:

### 1. Clone este repositório e entre no seu diretório:

```bash
git clone https://github.com/ufla-dsl/aula-git-parte-2.git
cd aula-git-parte-2
````

### 2. Execute o exemplo interativo passando um número como argumento:

```bash
go run ./app 100
```

### Saída esperada:

```
100°C = 212°F
100°F = 37.7778°C
```

Esse exemplo mostra a conversão do valor inserido para as duas escalas: Celsius e Fahrenheit.

### Código do exemplo (`app/main.go`):

```go
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ufla-gcc259/aula-git-parte-2/tempconv"
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
```

---

## Exemplo alternativo

Você também pode rodar o exemplo direto no `main.go`:

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

---

## Outras linguagens?

Versões da biblioteca *tempconv* para outras linguagens:

> *Todo*

---

## Licença

> *Todo*

---

## Como contribuir?

Escolha uma *issue* dentre as [disponíveis](https://github.com/ufla-gcc259/aula-git-parte-2/issues), avise à comunidade que você está trabalhando nela e envie um *Pull Request* quando terminar.

```

Se quiser, posso salvar isso em um arquivo ou te ajudar a criar a pasta `app/` com o `main.go`. Deseja que eu gere isso para você?
```
