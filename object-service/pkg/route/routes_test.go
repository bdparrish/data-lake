package route

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"

	"github.com/codingexplorations/data-lake/common/test/utilities"
	"github.com/codingexplorations/data-lake/object-service/pkg/controller"
	"github.com/codingexplorations/data-lake/object-service/test/mocks/repository"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	logger := utilities.TestLog{}
	router := gin.Default()
	repository := repository.NewObjectRepository(t)
	controller := controller.NewObjectController(&logger, repository)

	RegisterRoutes(router, *controller)

	routeMap := lo.Associate(
		router.Routes(),
		func(item gin.RouteInfo) (string, gin.HandlerFunc) {
			return fmt.Sprintf("%s %s", item.Method, item.Path), item.HandlerFunc
		},
	)

	assert.NotNil(t, routeMap, "failed to set routes for api")
	assert.Len(t, routeMap, 2)

	assert.Contains(t, routeMap, "GET /api/object/")
	assert.Equal(t, getHandlerFuncName(controller.GetObjects), getHandlerFuncName(routeMap["GET /api/object/"]))

	assert.Contains(t, routeMap, "GET /docs/*any", "failed to retrieve Swagger docs route")
}

func getHandlerFuncName(function gin.HandlerFunc) string {
	return runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name()
}
