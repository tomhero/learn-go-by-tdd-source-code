package main

import "fmt"

const spanish = "Spanish"
const french = "French"

const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "

const defalutEngSuffix = "World"

// Hello world function ja.
func Hello(name string, lang string) string {
	if name == "" {
		name = defalutEngSuffix
	}

	return greetingPrefix(lang) + name
}

// return whatever it's set to by just calling return rather than return prefix.
// signature we have made a named return value (prefix string).
func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishPrefix
		break
	case french:
		prefix = frenchPrefix
		break
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Tom", "French"))
}
