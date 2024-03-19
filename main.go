package main

import (
	"github.com/codeexplorations/data-lake/config"
	"github.com/codeexplorations/data-lake/pkg"
)

// main function that processes a local file
func main() {
	conf := config.GetConfig()

	pkg.ProcessFolder(conf.DataFolder)
}
