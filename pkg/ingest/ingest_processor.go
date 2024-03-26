package ingest

import (
	"fmt"
	"os"

	"github.com/bufbuild/protovalidate-go"
	models_v1 "github.com/codingexplorations/data-lake/models/v1"
	"github.com/codingexplorations/data-lake/pkg/config"
)

type IngestProcessor interface {
	ProcessFolder(folder string) ([]*models_v1.Object, error)
	ProcessFile(fileName string) (*models_v1.Object, error)
}

func GetIngestProcessor(conf *config.Config) IngestProcessor {
	switch os.Getenv(conf.IngestProcessorType) {
	case "local":
		return &LocalIngestProcessorImpl{}
	default:
		return &LocalIngestProcessorImpl{}
	}
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
