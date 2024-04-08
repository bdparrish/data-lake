package main

import (
	"fmt"
	golog "log"
	"os"
	"time"

	"github.com/codingexplorations/data-lake/pkg"
	"github.com/codingexplorations/data-lake/pkg/config"
	"github.com/codingexplorations/data-lake/pkg/db"
	"github.com/codingexplorations/data-lake/pkg/ingest"
	"github.com/codingexplorations/data-lake/pkg/log"
)

// main function that processes a local file
func main() {
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		migrate()
	} else {
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
}

func migrate() {
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
