package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/codingexplorations/data-lake/common/pkg/config"
	"github.com/codingexplorations/data-lake/common/pkg/models/v1/db"
	"github.com/codingexplorations/data-lake/common/pkg/models/v1/proto"
	"github.com/codingexplorations/data-lake/common/test/utilities"
	"github.com/codingexplorations/data-lake/object-service/test/mocks/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/encoding/protojson"
)

func initMocks(t *testing.T) (
	*utilities.TestLog,
	*config.Config,
	*repository.ObjectRepository[db.Object],
	ObjectController,
) {
	gin.SetMode(gin.TestMode)
	logger := utilities.TestLog{}
	conf := config.GetConfig()
	repo := repository.NewObjectRepository(t)

	controller := NewObjectController(&logger, repo)
	return &logger, conf, repo, *controller
}

func TestNewUserController(t *testing.T) {
	_, _, _, controller := initMocks(t)

	assert.NotNil(t, controller, "failed to construct user controller")
}

func TestObjectsController_GetObjects(t *testing.T) {
	_, _, repo, controller := initMocks(t)

	objects := []*db.Object{
		{
			Id:           uuid.New(),
			FileName:     "test 1",
			FileLocation: "test 1",
			ContentType:  "test 1",
			ContentSize:  1,
		},
		{
			Id:           uuid.New(),
			FileName:     "test 2",
			FileLocation: "test 2",
			ContentType:  "test 2",
			ContentSize:  2,
		},
	}

	repo.On("GetAll", 1, 2).Return(objects, nil)

	req, _ := http.NewRequest("GET", "/?page=1&size=2", nil)

	result := httptest.NewRecorder()

	context, _ := gin.CreateTestContext(result)
	context.Request = req

	controller.GetObjects(context)

	assert.Equal(t, http.StatusOK, result.Code)

	var response proto.ObjectGetAllResponse
	err := protojson.Unmarshal(result.Body.Bytes(), &response)
	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, response.Objects)
	assert.Len(t, response.Objects, 2)
	assert.Equal(t, objects[0].Id.String(), response.Objects[0].Id)
	assert.Equal(t, objects[0].FileName, response.Objects[0].FileName)
	assert.Equal(t, objects[0].FileLocation, response.Objects[0].FileLocation)
	assert.Equal(t, objects[0].ContentType, response.Objects[0].ContentType)
	assert.Equal(t, objects[0].ContentSize, response.Objects[0].ContentSize)
	assert.Equal(t, objects[1].Id.String(), response.Objects[1].Id)
	assert.Equal(t, objects[1].FileName, response.Objects[1].FileName)
	assert.Equal(t, objects[1].FileLocation, response.Objects[1].FileLocation)
	assert.Equal(t, objects[1].ContentType, response.Objects[1].ContentType)
	assert.Equal(t, objects[1].ContentSize, response.Objects[1].ContentSize)
}

func TestObjectsController_GetObjects_Empty(t *testing.T) {
	_, _, repo, controller := initMocks(t)

	repo.On("GetAll", 1, 10).Return([]*db.Object{}, nil)

	req, _ := http.NewRequest("GET", "/", nil)

	result := httptest.NewRecorder()

	context, _ := gin.CreateTestContext(result)
	context.Request = req

	controller.GetObjects(context)

	assert.Equal(t, http.StatusOK, result.Code)

	var response proto.ObjectGetAllResponse

	if result.Body.Len() != 0 {
		err := protojson.Unmarshal(result.Body.Bytes(), &response)
		if err != nil {
			t.Error(err)
		}
	}

	assert.Empty(t, response.Objects)
}
