package main

import (
	"github.com/codeexplorations/data-lake/pkg"
	"github.com/codeexplorations/data-lake/pkg/config"
)

// main function that processes a local file
func main() {
	conf := config.GetConfig()

	pkg.ProcessFolder(conf.DataFolder)
}
