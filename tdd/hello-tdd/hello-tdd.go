package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello ...
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Portuguese" {
		return "Olá, " + name
	}

	if language == "Spanish" {
		return "Hola, " + name
	}

	if language == "French" {
		return "Bonjour, " + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Lucas", ""))
}
