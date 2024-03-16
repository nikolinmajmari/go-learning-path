package main

import "fmt"

const (
	french             = "french"
	spanish            = "spanish"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjur, "
)

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	return helloPrefix(lang) + name
}

func helloPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		return spanishHelloPrefix
	case french:
		return frenchHelloPrefix
	default:
		return englishHelloPrefix
	}
}

func main() {
	fmt.Println(Hello("World", "english"))
}
