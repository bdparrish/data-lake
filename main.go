package main

import (
	"fmt"
	"github.com/codeexplorations/data-lake/models"
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

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	l := fmt.Sprintf("file://%s", pwd+"/"+s)

	object := models.Object{
		Filename:     s,
		FileLocation: l,
		ContentType:  "text/plain",
		ContentSize:  fileSize,
	}

	fmt.Printf("Object: %+v\n", object)
}
