package main

import (
	"fmt"
)

const (
	french             = "French"
	spanish            = "Spanish"
	englishGreetPrefix = "Hello, "
	spanishGreetPrefix = "Hola, "
	frenchGreetPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchGreetPrefix
	case spanish:
		prefix = spanishGreetPrefix
	default:
		prefix = englishGreetPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Aryan", "Spanish"))
}
