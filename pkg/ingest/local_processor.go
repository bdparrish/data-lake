package ingest

import (
	"fmt"
	golog "log"
	"os"
	"strings"

	dbModels "github.com/codingexplorations/data-lake/models/v1/db"
	"github.com/codingexplorations/data-lake/pkg/config"
	"github.com/codingexplorations/data-lake/pkg/db"
	"github.com/codingexplorations/data-lake/pkg/log"
	"gorm.io/gorm"
)

type LocalIngestProcessorImpl struct {
	logger log.Logger
	db     *gorm.DB
}

func NewLocalIngestProcessor() *LocalIngestProcessorImpl {
	logger, err := log.GetLogger()
	if err != nil {
		golog.Fatalf("couldn't create logger: %v\n", err)
	}

	conf := config.GetConfig()

	return &LocalIngestProcessorImpl{
		logger: logger,
		db:     db.NewPostgresDb(conf),
	}
}

// ProcessFolder processes the file
func (processor *LocalIngestProcessorImpl) ProcessFolder(folder string) ([]*dbModels.Object, error) {
	entries, err := os.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	processedObjects := make([]*dbModels.Object, 0)

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

	os.Remove(folder)

	return processedObjects, nil
}

// ProcessFile processes the file
func (processor *LocalIngestProcessorImpl) ProcessFile(fileName string) (*dbModels.Object, error) {
	// read the file
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	fileSize := len(data)

	pathSplit := strings.Split(fileName, "/")

	object := &dbModels.Object{
		FileName:     pathSplit[len(pathSplit)-1],
		FileLocation: fileName,
		ContentType:  "text/plain",
		ContentSize:  int32(fileSize),
	}

	valid, err := validate(object)
	if err != nil {
		processor.logger.Error(fmt.Sprintf("error validating object: %v\n", err))
		return nil, err
	}

	if !valid {
		processor.logger.Error(fmt.Sprintf("object is not valid: %v\n", object))
		return nil, nil
	}

	os.Remove(fileName)

	if err := processor.db.Create(&object).Error; err != nil {
		processor.logger.Error(fmt.Sprintf("error creating object record - %v", err))
		return nil, err
	}

	return object, nil
}
