package main

import "fmt"

const englishPrefix = "Hello, "
const defalutEngSuffix = "World"

// Hello world function ja.
func Hello(name string) string {
	if name != "" {
		return englishPrefix + name
	}
	return englishPrefix + defalutEngSuffix
}

func main() {
	fmt.Println(Hello("Tom"))
}
