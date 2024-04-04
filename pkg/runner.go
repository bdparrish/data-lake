package pkg

import (
	"time"

	"github.com/codingexplorations/data-lake/pkg/config"
	"github.com/codingexplorations/data-lake/pkg/ingest"
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
	r.Config.Print()
	for {
		r.Processor.ProcessFolder(r.Config.DataFolder)
		time.Sleep(10 * time.Second)
	}
}
