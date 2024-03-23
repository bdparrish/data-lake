package main

import (
	"github.com/codeexplorations/data-lake/config"
	"github.com/codeexplorations/data-lake/pkg"
	"github.com/codeexplorations/data-lake/pkg/ingest"
)

// main function that processes a local file
func main() {
	conf := config.GetConfig()
	processor := ingest.GetIngestProcessor(conf)

	pkg.NewRunner(conf, processor).Run()
}
