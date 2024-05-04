package main

import (
	"fmt"
	golog "log"
	"os"

	"github.com/codingexplorations/data-lake/common/pkg/config"
	"github.com/codingexplorations/data-lake/common/pkg/db"
	"github.com/codingexplorations/data-lake/common/pkg/log"
)

// main function that processes a local file
func main() {
	config := config.GetConfig()
	logger, err := log.GetLogger()
	if err != nil {
		golog.Println("failed to get logger")
		golog.Printf("error: %v", err)
		os.Exit(1)
	}

	logger.Info("Starting DB migration ...")
	dbMigrator, err := db.NewMigrator(logger, config)
	if err != nil {
		logger.Error(fmt.Sprintf("error in initializing DB migrator: %v", err))
		os.Exit(2)
	}
	err = dbMigrator.Migrate()
	if err != nil {
		logger.Error(fmt.Sprintf("error in migrating DB: %v", err))
		os.Exit(3)
	}
}
