package greetings

import "fmt"

// Hello returns a greeting for a named person
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome", name)

	return message
}