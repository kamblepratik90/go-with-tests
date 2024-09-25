package main

import "fmt"

// func Hello() string {
// 	return "Hello World"
// }

const englishHelloPrefix = "Hello, "

func Hello(msg string) string {

	if msg == "" {
		return englishHelloPrefix + "World"
		// msg = "World"
	}
	return englishHelloPrefix + msg
}

func main() {
	fmt.Println(Hello("OS"))
}
