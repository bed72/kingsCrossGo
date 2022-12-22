package utils

import (
	"fmt"
	"os"
)

func ConnectionURLBuilder(str string) (string, error) {
	var url string

	switch str {
	case "fiber":
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("SERVER_HOST"),
			os.Getenv("SERVER_PORT"),
		)
	default:
		return "", fmt.Errorf("connection name '%v' is not supported", str)
	}

	return url, nil
}
