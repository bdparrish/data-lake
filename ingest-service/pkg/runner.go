package pkg

import (
	"log"

	"github.com/codingexplorations/data-lake/common/pkg/config"
	"github.com/codingexplorations/data-lake/ingest-service/pkg/ingest"
)

type Runner struct {
	Config    *config.Config
	Processor ingest.IngestProcessor
}

func NewRunner(conf *config.Config, processor ingest.IngestProcessor) *Runner {
	return &Runner{
		Config:    conf,
		Processor: processor,
	}
}

func (r *Runner) Run() {
	_, err := r.Processor.ProcessFolder(r.Config.DataFolder)
	if err != nil {
		log.Printf("couldn't process folder %v: %v\n", r.Config.DataFolder, err)
	}
}
