package ingest

import (
	"fmt"
	"log"

	"github.com/bufbuild/protovalidate-go"
	models_v1 "github.com/codingexplorations/data-lake/models/v1"
	"github.com/codingexplorations/data-lake/pkg/config"
)

type IngestProcessor interface {
	ProcessFolder(folder string) ([]*models_v1.Object, error)
	ProcessFile(fileName string) (*models_v1.Object, error)
}

func GetIngestProcessor(conf *config.Config) IngestProcessor {
	log.Println("here")
	switch conf.IngestProcessorType {
	case "local":
		log.Println("Using local ingest processor")
		return NewLocalIngestProcessor()
	case "localstack":
		log.Println("Using localstack ingest processor")
		return NewS3IngestProcessorImpl(conf)
	default:
		log.Println("Using default ingest processor")
		return NewLocalIngestProcessor()
	}
}

func validate(object *models_v1.Object) (bool, error) {
	validator, err := protovalidate.New()
	if err != nil {
		return false, fmt.Errorf("failed to initialize proto validator: %v", err)
	}

	if err := validator.Validate(object); err != nil {
		return false, fmt.Errorf("failed to validate object: %v", err)
	}

	return true, nil
}
