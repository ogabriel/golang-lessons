package main

import "fmt"
import "reflect"

func main() {
	exibeNomes()
}

func exibeNomes() {
	nomes := []string{"Douglas", "Daniel", "Bernardo"}

	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Printf("O meu slice tem: ", len(nomes), "capacidade: ", cap(nomes))

	// dobra a capacidade do array, igual rust
	nomes = append(nomes, "Aparecida")

	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Printf("O meu slice tem: ", len(nomes), "capacidade: ", cap(nomes))

}
