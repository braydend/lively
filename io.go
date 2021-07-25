package main

import (
	"io"
	"os"
)

func openFile(filename string) *os.File {
	data, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	return data
}

func ReadFile(filename string) []byte {
	file := openFile(filename)
	data, err := io.ReadAll(file)

	if err != nil {
		panic(err)
	}

	return data
}
