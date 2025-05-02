*tempconv*
=====
Biblioteca escrita em Go para conversões simples de temperatura. Extraída do livro **A Linguagem de Programação Go**, de Alan A. A. Donovan e Brian Kernighan. 

Como usar?
----
Para rodar o programa localmente e entender o que ele faz, siga os passos abaixo:

1. Clone este repositório, entre no seu diretório e rode o programa

```shell
git clone https://github.com/ufla-dsl/aula-git-parte-2.git
cd aula-git-parte-2
cd go
go run main.go
```

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

### Python:

```shell
cd python
python main.py
```

### Java:

```shell
cd java
javac TempConv.java
java TempConv
```

### C:

```shell
cd c
gcc tempconv.c -o tempconv
./tempconv
```

### C++:
 
```shell
cd cpp
g++ tempconv.cpp -o tempconv
./tempconv
```

### Typescript:

```shell
cd typescript
ts-node main.ts
```

Licença
-----

> *Todo*


Como contribuir?
----
Escolha uma *issue* dentre as [disponíveis](https://github.com/ufla-gcc259/aula-git-parte-2/issues), avise à comunidade que você está trabalhando nela e envie um *Pull Request*, quando terminar.
