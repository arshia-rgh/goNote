package utils

import (
	"fmt"
	"os"
)

func CheckError(err error, message string) {
	if err != nil {
		fmt.Printf("%s: %v\n", message, err)
		os.Exit(1)
	}
}
