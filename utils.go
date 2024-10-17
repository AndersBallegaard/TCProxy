package main

import (
	"log"
	"os"
)

func checkErrFatal(err any) {
	if err != nil {
		log.Fatalln(err)
	}
}

func readFileByte(filepath string) []byte {
	data, err := os.ReadFile(filepath)
	checkErrFatal(err)
	return data
}
