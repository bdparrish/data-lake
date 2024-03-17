package main

import (
	"fmt"
	"github.com/codeexplorations/data-lake/models/proto"
	"os"
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

	object := proto.Object{
		FileName:     fileName,
		FileLocation: l,
		ContentType:  "text/plain",
		ContentSize:  int64(fileSize),
	}

	fmt.Printf("Object: %+v\n", object)
}
