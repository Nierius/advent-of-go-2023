package utils

import (
	"log"
	"os"
)

func ReadEntireFileToString(fileName string) string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return string(file)
}
