package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/codeexplorations/data-lake/config"
	models_v1 "github.com/codeexplorations/data-lake/models/v1"
)

// main function that processes a local file
func main() {
	conf := config.GetConfig()

	ProcessFolder(conf.DataFolder)
}

// ProcessFile processes the file
func ProcessFolder(folder string) {
	entries, err := os.ReadDir(folder)
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			ProcessFolder(folder + "/" + entry.Name())
		} else {
			ProcessFile(folder + "/" + entry.Name())
		}
	}
}

// ProcessFile processes the file
func ProcessFile(fileName string) {
	// read the file
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	fileSize := len(data)

	pathSplit := strings.Split(fileName, "/")

	object := models_v1.Object{
		FileName:     pathSplit[len(pathSplit)-1],
		FileLocation: fileName,
		ContentType:  "text/plain",
		ContentSize:  int64(fileSize),
	}

	fmt.Printf("Object: %+v\n", &object)
}
