package utils

import (
	"log"
	"os"
)

func ReadFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Println("Error: ", err)
		os.Exit(1)
	}

	return string(data)
}
