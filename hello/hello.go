package main

import (
	"fmt"
)

func main() {
	fmt.Println(Hello("Frankie", "English"))
}

const (
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPredix = "Hola, "
	frenchHelloPrefix  = "Bonjur, "
)

func Hello(name string, language string) string {
	if len(name) == 0 {
		name = "World"
	}

	prefix := getPrefix(language)

	return prefix + name
}

func getPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPredix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
