package route

import (
	"github.com/codingexplorations/data-lake/object-service/pkg/controller"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(
	router *gin.Engine,
	controller controller.ObjectController,
) {
	api := router.Group("/api/object")
	{
		api.GET("/", controller.GetObjects)
	}

	router.GET(
		"/docs/*any",
		ginSwagger.WrapHandler(
			swaggerFiles.Handler,
			ginSwagger.DefaultModelsExpandDepth(1),
			ginSwagger.DocExpansion("list"),
			ginSwagger.DeepLinking(true),
		),
	)
}
