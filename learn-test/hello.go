package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanish = "Spanish"
const french = "French"
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func main() {
	fmt.Println(Hello("world", spanish))
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) string {
	switch language {
	case spanish:
		return spanishHelloPrefix
	case french:
		return frenchHelloPrefix
	}
	return englishHelloPrefix

}
