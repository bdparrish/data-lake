package ingest

import (
	"fmt"
	"os"
	"strings"

	"github.com/bufbuild/protovalidate-go"
	models_v1 "github.com/codeexplorations/data-lake/models/v1"
)

type LocalIngestProcessorImpl struct{}

// ProcessFile processes the file
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

	validate(object)

	return object, nil
}

func validate(object *models_v1.Object) (bool, error) {
	validator, err := protovalidate.New()
	if err != nil {
		return false, fmt.Errorf("failed to initialize proto validator: %v", err)
	}

	if err := validator.Validate(object); err != nil {
		return false, fmt.Errorf("failed to validate object: %v", err)
	} else {
		fmt.Println("object is valid")
	}

	return true, nil
}
