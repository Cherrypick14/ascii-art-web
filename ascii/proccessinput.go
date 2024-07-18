package ascii

import (
	"strings"
)

func Processinput(inputText string) []string {
	// Split our input text to a string slice and separate with a newline.
	words := strings.Split(inputText, "\n")

	return words
}
