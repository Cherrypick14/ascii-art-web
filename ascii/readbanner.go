package ascii

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func ReadBannerFile(banner string) ([]string, error) {
	path := "banners/"
	if banner == "" {
		banner = path + "standard.txt"
	}

	fileName := path + banner + ".txt"
	fmt.Println(fileName)

	hash := checksum([]byte(fileName))

	hashStandard := "8c8578d2764505485aa7e98258df22b5b2be2fd2722ee918076c72b2df0974e4"
	hashShadow := "974747921efbff0708af084ab495eeb471a6743fd7d1a02bf2bb5a0179a520fb"
	hashThinkertoy := "f8ac20c65b51ca7640958022a063ec46572637960c5a821a24b197d38f75afb1"

	if hash != hashStandard && hash != hashShadow && hash != hashThinkertoy {
		return nil, fmt.Errorf("error reading file")
	}

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

func checksum(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}
