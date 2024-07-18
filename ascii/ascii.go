package ascii

import (
	// "fmt"
	"strings"
)

func AsciiArt(words []string, contents2 []string) (string, error) {
	var result strings.Builder

	countSpace := 0

	for _, word := range words {
		if word != "" {
			for i := 0; i < 8; i++ {
				for _, char := range word {
					if char == '\n' {
						continue
					}
					if !(char >= 32 && char <= 126) {
						return "Error: Input contains non-ASCII characters", nil
					}
					index := int(char-' ')*9 + 1 + i
					result.WriteString(string(contents2[index]))
				}
				result.WriteString("\n")
			}
		} else {
			countSpace++
			if countSpace < len(words) {
				result.WriteString("\n")
			}
		}
	}
	return result.String(), nil
}
