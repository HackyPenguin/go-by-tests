package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const germanHelloPrefix = "Hallo, "

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "french":
		prefix = frenchHelloPrefix
	case "spanish":
		prefix = spanishHelloPrefix
	case "german":
		prefix = germanHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name

}

func main() {
	fmt.Println(Hello("Barry", "spanish"))
}
