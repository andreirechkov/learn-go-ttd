package iteration

import "strings"

func Repeat(character string, count int) string {
	var result strings.Builder
	for i := 0; i < count; i++ {
		result.WriteString(character)
	}
	return result.String()
}