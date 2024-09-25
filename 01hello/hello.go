package main

import "fmt"

// func Hello() string {
// 	return "Hello World"
// }

// const Spanish = "Spanish"
// const French = "French"
// const englishHelloPrefix = "Hello, "
// const SpanishHelloPrefix = "Hola, "
// const FrenchHelloPrefix = "Bonjour, "

const (
	Spanish            = "Spanish"
	French             = "French"
	englishHelloPrefix = "Hello, "
	SpanishHelloPrefix = "Hola, "
	FrenchHelloPrefix  = "Bonjour, "
)

func Hello(msg string, language string) string {

	if msg == "" {
		return englishHelloPrefix + "World"
		// msg = "World"
	}
	// if language == Spanish {
	// 	return SpanishHelloPrefix + msg
	// }
	// if language == French {
	// 	return FrenchHelloPrefix + msg
	// }
	// return englishHelloPrefix + msg

	// -------------------------------

	// langPrefix := englishHelloPrefix

	// switch language {
	// case Spanish:
	// 	langPrefix = SpanishHelloPrefix
	// case French:
	// 	langPrefix = FrenchHelloPrefix
	// }

	// return langPrefix + msg

	// -----------------------

	return greetingPrefix(language) + msg

}

func greetingPrefix(language string) (langPrefix string) {
	switch language {
	case Spanish:
		langPrefix = SpanishHelloPrefix
	case French:
		langPrefix = FrenchHelloPrefix
	default:
		langPrefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("OS", ""))
}
