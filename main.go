package main

import (
	"github.com/codingexplorations/data-lake/config"
	"github.com/codingexplorations/data-lake/pkg"
	"github.com/codingexplorations/data-lake/pkg/ingest"
)

// main function that processes a local file
func main() {
	conf := config.GetConfig()
	processor := ingest.GetIngestProcessor(conf)

	pkg.NewRunner(conf, processor).Run()
}
