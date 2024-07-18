package ascii

import (
	"bufio"
	"os"
)

func ReadBannerFile(banner string) ([]string, error) {
	path := "ascii/"
	if banner == "" {
		banner = path + "standard.txt"
	}
	fileName := path + banner + ".txt"
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var contents []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}

	return contents, scanner.Err()
}
