package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/codingexplorations/data-lake/common/pkg/config"
	"github.com/codingexplorations/data-lake/common/pkg/log"
	"github.com/codingexplorations/data-lake/common/pkg/models/v1/db"
	"github.com/codingexplorations/data-lake/common/test/utilities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func initMocks() (sqlmock.Sqlmock, ObjectRepository[db.Object]) {
	logger := utilities.TestLog{}
	conf := config.GetConfig()

	mockDB, gormDB := utilities.NewMockDB()

	repo := NewObjectRepository(&logger, conf, gormDB)
	return mockDB, repo
}

func TestNewObjectRepository(t *testing.T) {
	logger, _ := log.GetLogger()
	conf := config.GetConfig()
	_, gormDB := utilities.NewMockDB()

	repo := NewObjectRepository(logger, conf, gormDB)

	assert.NotNil(t, repo, "failed to create a new repository")
}

func TestObjectRepositoryImpl_GetAll(t *testing.T) {
	mockDb, repo := initMocks()

	object1 := &db.Object{
		Id:           uuid.New(),
		FileName:     "testing name",
		FileLocation: "testing/location/directory",
		ContentType:  "application/json",
		ContentSize:  1234,
	}
	object2 := &db.Object{
		Id:           uuid.New(),
		FileName:     "testing name 2",
		FileLocation: "testing/location/directory2",
		ContentType:  "application/json",
		ContentSize:  5678,
	}

	sqlRows := sqlmock.NewRows([]string{"id", "file_name", "file_location", "content_type", "content_size"}).
		AddRow(object1.Id, object1.FileName, object1.FileLocation, object1.ContentType, object1.ContentSize).
		AddRow(object2.Id, object2.FileName, object2.FileLocation, object2.ContentType, object2.ContentSize)

	mockDb.ExpectQuery("^SELECT .* FROM .*").WillReturnRows(sqlRows)

	objects, err := repo.GetAll(1, 10)
	if err != nil {
		t.Error("failed to get object from the data store")
	}

	assert.NotNil(t, objects[0].Id)
	assert.Equal(t, "testing name", objects[0].FileName)
	assert.Equal(t, "testing/location/directory", objects[0].FileLocation)
	assert.Equal(t, "application/json", objects[0].ContentType)
	assert.Equal(t, int32(1234), objects[0].ContentSize)
	assert.NotNil(t, objects[1].Id)
	assert.Equal(t, "testing name 2", objects[1].FileName)
	assert.Equal(t, "testing/location/directory2", objects[1].FileLocation)
	assert.Equal(t, "application/json", objects[1].ContentType)
	assert.Equal(t, int32(5678), objects[1].ContentSize)
}

func TestObjectRepositoryImpl_Get(t *testing.T) {
	mockDb, repo := initMocks()

	object := &db.Object{
		Id:           uuid.New(),
		FileName:     "testing name",
		FileLocation: "testing/location/directory",
		ContentType:  "application/json",
		ContentSize:  1234,
	}

	sqlRows := sqlmock.NewRows([]string{"id", "file_name", "file_location", "content_type", "content_size"}).
		AddRow(object.Id, object.FileName, object.FileLocation, object.ContentType, object.ContentSize)

	mockDb.ExpectQuery("^SELECT .* WHERE id .*").WillReturnRows(sqlRows)

	response, err := repo.Get(object.Id)
	if err != nil {
		t.Error("failed to get object from the data store")
	}

	assert.NotNil(t, response.Id)
	assert.Equal(t, "testing name", response.FileName)
	assert.Equal(t, "testing/location/directory", response.FileLocation)
	assert.Equal(t, "application/json", response.ContentType)
	assert.Equal(t, int32(1234), response.ContentSize)
}
