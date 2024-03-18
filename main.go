package main

import (
	"fmt"
	"os"

	models_v1 "github.com/codeexplorations/data-lake/models/v1"
)

// main function that processes a local file
func main() {
	ProcessFile("sample.txt")
}

// ProcessFile processes the file
func ProcessFile(fileName string) {
	// read the file
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	fileSize := len(data)

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	l := fmt.Sprintf("file://%s", pwd+"/"+fileName)

	object := models_v1.Object{
		FileName:     fileName,
		FileLocation: l,
		ContentType:  "text/plain",
		ContentSize:  int64(fileSize),
	}

	fmt.Printf("Object: %+v\n", object)
}
