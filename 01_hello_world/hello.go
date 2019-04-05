package main

import "fmt"

// Hello world function ja.
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("Tom"))
}
