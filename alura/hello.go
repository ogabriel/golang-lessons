package main

import "fmt"
import "reflect"

func main() {
	// da pra declarar váriaveis de várias formas diferentes:
	var nome = "Gabriel"
	var idade int = 21
	versao := 1.1

	fmt.Println("Olá, sr.", nome, "a sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("O tipo da variável versao é", reflect.TypeOf(versao))
}
