package main

import (
	"fmt"
	"os"

	"github.com/codingexplorations/data-lake/common/pkg/config"
	"github.com/codingexplorations/data-lake/common/pkg/db"
	"github.com/codingexplorations/data-lake/common/pkg/log"
	"github.com/codingexplorations/data-lake/object-service/pkg/controller"
	"github.com/codingexplorations/data-lake/object-service/pkg/repository"
	"github.com/codingexplorations/data-lake/object-service/pkg/route"
	"github.com/gin-gonic/gin"
)

// @title			Object Service
// @version		0.0.1
// @description	Object service for core object operations.
// @termsOfService	http://swagger.io/terms/
// @BasePath		/api/object/
func main() {
	conf := config.GetConfig()

	logger, err := log.GetLogger()
	if err != nil {
		panic("failed to get logger")
	}

	store := db.NewPostgresDb(conf)

	repo := repository.NewObjectRepository(logger, conf, store)

	objectController := controller.NewObjectController(logger, repo)

	engine := gin.New()
	route.RegisterRoutes(engine, *objectController)
	engine.Use(gin.Recovery())

	err = engine.Run(fmt.Sprintf(":%d", 8000))
	if err != nil {
		logger.Error(fmt.Sprintf("error in starting gin engine: %v", err.Error()))
		os.Exit(2)
	}
}
