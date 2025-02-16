package main

import "fmt"

const (
	spanish = "Spanish"
	french = "French"
	
	englishHelloPrefix = "Hello, "
	spanisHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
)

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	return gettingPrefix(lang) + name
}

func gettingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanisHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Andrey", ""))
}