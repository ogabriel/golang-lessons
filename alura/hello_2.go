package main

import "fmt"

func main() {
	// da pra declarar váriaveis de várias formas diferentes:
	nome, idade := devolveNomeEIdade()

	fmt.Println("Olá, sr.", nome, "a sua idade é", idade)
}

func devolveNomeEIdade() (string, int) {
	nome := "Gabriel"
	idade := 21
	return nome, idade

}
