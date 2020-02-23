package main

import "fmt"

const englishHelloPrefix = "Hello, "
const portugueseHelloPrefix = "Olá, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func getGreetingPrefix(language string) string {
	prefix := englishHelloPrefix

	switch language {
	case "Portuguese":
		prefix = portugueseHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	}

	return prefix
}

// Hello ...
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := getGreetingPrefix(language)

	return prefix + name
}

func main() {
	fmt.Println(Hello("Lucas", ""))
}
