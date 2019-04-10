package iteration

// Repeat function returns the character n times
func Repeat(character string, count int) string {
	var result string
	for i := 0; i < count; i++ {
		result += character
	}
	return result
}
