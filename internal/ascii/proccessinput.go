package ascii

import (
	"fmt"
	"os"
	"strings"
)

func Processinput(text string) []string {
	// Grab string to generate Ascii represatantion.
	inputText := os.Args[1]

	switch inputText {
	case "":
		return nil
	case "\\a", "\\0", "\\f", "\\v", "\\r":
		fmt.Println("Error: Non printable character", inputText)
		return nil
	}

	inputText = strings.ReplaceAll(inputText, "\\t", "    ")
	inputText = strings.ReplaceAll(inputText, "\\b", "\b")
	inputText = strings.ReplaceAll(inputText, "\\n", "\n")
	// Logic process for handlng the backspace.
	for i := 0; i < len(inputText); i++ {
		indexB := strings.Index(inputText, "\b")
		if indexB > 0 {
			inputText = inputText[:indexB-1] + inputText[indexB+1:]
		}
	}
	// Split our input text to a string slice and separate with a newline.
	words := strings.Split(inputText, "\n")

	return words
}
