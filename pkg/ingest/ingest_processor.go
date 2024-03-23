package ingest

import (
	"os"

	"github.com/codeexplorations/data-lake/config"
	models_v1 "github.com/codeexplorations/data-lake/models/v1"
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
