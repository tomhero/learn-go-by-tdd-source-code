package main

import "fmt"

const spanish = "Spanish"

const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "

const defalutEngSuffix = "World"

// Hello world function ja.
func Hello(name string, lang string) string {
	if name == "" {
		name = defalutEngSuffix
	}

	if lang == spanish {
		return spanishPrefix + name
	}

	return englishPrefix + name
}

func main() {
	fmt.Println(Hello("Tom", ""))
}
