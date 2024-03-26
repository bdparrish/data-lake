package ingest

import (
	"os"

	models_v1 "github.com/codeexplorations/data-lake/models/v1"
	"github.com/codeexplorations/data-lake/pkg/config"
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
