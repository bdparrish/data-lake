package main

import (
	"fmt"
	"os"
)

// main function that processes a local file
func main() {
	ProcessFile("sample.txt")
}

// ProcessFile processes the file
func ProcessFile(s string) {
	// read the file
	data, err := os.ReadFile(s)
	if err != nil {
		panic(err)
	}

	fileSize := len(data)

	fmt.Printf("File size: %d\n", fileSize)
}
