package main

import (
	"os"
	"time"

	"github.com/codingexplorations/data-lake/common/pkg/config"
	"github.com/codingexplorations/data-lake/common/pkg/log"
	"github.com/codingexplorations/data-lake/ingest-service/pkg"
	"github.com/codingexplorations/data-lake/ingest-service/pkg/ingest"
)

// main function that processes a local file
func main() {
	logger := log.NewConsoleLog()

	for _, e := range os.Environ() {
		logger.Info(e)
	}

	conf := config.GetConfig()
	processor := ingest.GetIngestProcessor(conf)

	r := pkg.NewRunner(conf, processor)

	r.Config.Print()

	for {
		r.Run()
		time.Sleep(10 * time.Second)
	}
}
