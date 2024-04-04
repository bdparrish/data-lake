package main

import (
	"os"

	"github.com/codingexplorations/data-lake/pkg"
	"github.com/codingexplorations/data-lake/pkg/config"
	"github.com/codingexplorations/data-lake/pkg/ingest"
	"github.com/codingexplorations/data-lake/pkg/log"
)

// main function that processes a local file
func main() {
	logger := log.NewConsoleLog()

	for _, e := range os.Environ() {
		// pair := strings.SplitN(e, "=", 2)
		logger.Info(e)
	}

	conf := config.GetConfig()
	processor := ingest.GetIngestProcessor(conf)

	pkg.NewRunner(conf, processor).Run()
}
