package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserData(prompt string) (string, error) {
	fmt.Println(prompt)

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil

}
