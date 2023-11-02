package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Gabriel")
	fmt.Println(message)
}
