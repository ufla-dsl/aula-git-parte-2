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

Outras linguagens?
----
Versões da biblioteca *tempconv* para outras linguagens:

- Java

Para rodar o programa localmente e entender o que ele faz, siga os passos abaixo:

1. Clone o [repositório](https://github.com/WillianBrandao/aula-git-parte-2), entre no seu diretório e rode o programa

```bash
git clone https://github.com/WillianBrandao/aula-git-parte-2.git
```

```bash 
cd aula-git-parte-2 
```
Compile os arquivos  .java

```bash
javac Main.java tempconv/TempConv.java
```

Rode o programa

```bash
java Main
```


> *Todo*


Licença
-----

> *Todo*


Como contribuir?
----
Escolha uma *issue* dentre as [disponíveis](https://github.com/ufla-gcc259/aula-git-parte-2/issues), avise à comunidade que você está trabalhando nela e envie um *Pull Request*, quando terminar.
