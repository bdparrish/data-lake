package main

import (
	"github.com/codeexplorations/data-lake/config"
)

// main function that processes a local file
func main() {
	conf := config.GetConfig()

	ProcessFolder(conf.DataFolder)
}
