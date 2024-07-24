package ascii

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AsciiArt(input string, banner string) (string, error) {
	var result strings.Builder

	input = strings.ReplaceAll(input, "\r", "\n")
	words := strings.Split(input, "\n")

	path := "ascii/banner/"

	banner2 := path + banner + ".txt"

	file, err := os.Open(banner2)
	if err != nil {
		return "404", err
	}

	defer file.Close()

	var contents []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}

	countSpace := 0

	for _, word := range words {
		if word != "" {
			for i := 0; i < 8; i++ {
				for _, char := range word {
					if char == '\n' {
						continue
					}
					if !(char >= 32 && char <= 126) {
						return "400", fmt.Errorf("error")
					}
					index := int(char-' ')*9 + 1 + i
					result.WriteString(string(contents[index]))
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
