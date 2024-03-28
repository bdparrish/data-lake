package ingest

import (
	"fmt"
	"os"
	"strings"

	models_v1 "github.com/codingexplorations/data-lake/models/v1"
	"github.com/codingexplorations/data-lake/pkg/log"
)

type LocalIngestProcessorImpl struct{}

// ProcessFolder processes the file
func (processor *LocalIngestProcessorImpl) ProcessFolder(folder string) ([]*models_v1.Object, error) {
	entries, err := os.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	processedObjects := make([]*models_v1.Object, 0)

	for _, entry := range entries {
		if entry.IsDir() {
			if processedFolderObjects, err := processor.ProcessFolder(folder + "/" + entry.Name()); err != nil {
				return nil, err
			} else {
				processedObjects = append(processedObjects, processedFolderObjects...)
			}
		} else {
			if processedFile, err := processor.ProcessFile(folder + "/" + entry.Name()); err != nil {
				return nil, err
			} else {
				processedObjects = append(processedObjects, processedFile)
			}
		}
	}

	return processedObjects, nil
}

// ProcessFile processes the file
func (processor *LocalIngestProcessorImpl) ProcessFile(fileName string) (*models_v1.Object, error) {
	// read the file
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	fileSize := len(data)

	pathSplit := strings.Split(fileName, "/")

	object := &models_v1.Object{
		FileName:     pathSplit[len(pathSplit)-1],
		FileLocation: fileName,
		ContentType:  "text/plain",
		ContentSize:  int32(fileSize),
	}

	valid, err := validate(object)
	if err != nil {
		log.NewConsoleLog().Error(fmt.Sprintf("error validating object: %v\n", err))
		return nil, err
	}

	if !valid {
		log.NewConsoleLog().Error(fmt.Sprintf("object is not valid: %v\n", object))
		return nil, nil
	}

	return object, nil
}
