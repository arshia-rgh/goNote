package utils

import "fmt"

func GetUserData(prompt string) (string, error) {
	fmt.Println(prompt)
	var input string
	_, err := fmt.Scan(&input)

	if err != nil {
		return "", err
	}

	return input, nil

}
