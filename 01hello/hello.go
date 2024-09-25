package main

import "fmt"

// func Hello() string {
// 	return "Hello World"
// }

func Hello(msg string) string {
	return "Hello " + msg
}

func main() {
	fmt.Println(Hello("OS"))
}
