package ingest

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/codingexplorations/data-lake/test/utilities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestFolderIngest_ProcessFolder_CheckDepth(t *testing.T) {
	tests := []struct {
		name      string
		folder    string
		subfolder string
		location  string
	}{
		{
			name:      "directory",
			folder:    "/test",
			subfolder: "files",
			location:  "/test.txt",
		},
		{
			name:     "file",
			folder:   "/test",
			location: "/test.txt",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockDb, gormDb := utilities.NewMockDB()

			pwd, _ := os.Getwd()
			fileName := filepath.Join(pwd, tc.folder, tc.subfolder, "test.txt")

			objectsRow := sqlmock.NewRows(
				[]string{"id", "file_name", "file_location", "content_type", "content_size"},
			).AddRow(
				uuid.New().String(), "test.txt", fileName, "text/plain", 15,
			)
			mockDb.ExpectQuery("^INSERT INTO \"objects\".*").WithArgs().WillReturnRows(objectsRow)

			processor := &LocalIngestProcessorImpl{}
			processor.db = gormDb

			_ = os.MkdirAll(pwd+tc.folder+"/"+tc.subfolder, os.ModePerm)

			d1 := []byte("This is a test.")
			_ = os.WriteFile(fileName, d1, 0644)

			processedObjects, err := processor.ProcessFolder(pwd + tc.folder)

			os.Remove(pwd + tc.folder + "/" + tc.subfolder)

			assert.Nil(t, err)
			assert.Len(t, processedObjects, 1)
			assert.NotEmpty(t, processedObjects[0].Id.String())
			assert.Equal(t, "test.txt", processedObjects[0].FileName)
			assert.Equal(t, filepath.Join(pwd, tc.folder, tc.subfolder, tc.location), processedObjects[0].FileLocation)
			assert.Equal(t, "text/plain", processedObjects[0].ContentType)
			assert.Equal(t, int32(15), processedObjects[0].ContentSize)
		})
	}
}

func TestFolderIngest_ProcessFile_Success(t *testing.T) {
	mockDb, gormDb := utilities.NewMockDB()

	pwd, _ := os.Getwd()
	fileName := filepath.Join(pwd, "test", "test.txt")

	objectsRow := sqlmock.NewRows(
		[]string{"id", "file_name", "file_location", "content_type", "content_size"},
	).AddRow(
		uuid.New().String(), "test.txt", fileName, "text/plain", 15,
	)
	mockDb.ExpectQuery("^INSERT INTO \"objects\".*").WithArgs().WillReturnRows(objectsRow)

	processor := &LocalIngestProcessorImpl{}
	processor.db = gormDb

	_ = os.Mkdir("test", os.ModePerm)

	d1 := []byte("test go")
	_ = os.WriteFile(fileName, d1, 0644)

	processedObject, err := processor.ProcessFile(fileName)

	os.Remove("test")

	assert.Nil(t, err)
	assert.NotEmpty(t, processedObject.Id.String())
	assert.Equal(t, "test.txt", processedObject.FileName)
	assert.Equal(t, fileName, processedObject.FileLocation)
	assert.Equal(t, "text/plain", processedObject.ContentType)
	assert.Equal(t, int32(15), processedObject.ContentSize)
}

func TestFolderIngest_ProcessFolder_Failure(t *testing.T) {
	processor := &LocalIngestProcessorImpl{}

	pwd, _ := os.Getwd()

	folder := pwd + "/../../test/files/missing"

	processedObject, err := processor.ProcessFolder(folder)

	assert.Error(t, err)
	assert.Equal(t, "open /Users/benjaminparrish/Development/Personal/data-lake/pkg/ingest/../../test/files/missing: no such file or directory", err.Error())
	assert.Nil(t, processedObject)
}

func TestFolderIngest_ProcessFile_Failure(t *testing.T) {
	processor := &LocalIngestProcessorImpl{}

	pwd, _ := os.Getwd()

	fileName := pwd + "/../../test/files/ingest/missing.txt"

	processedObject, err := processor.ProcessFile(fileName)

	assert.Error(t, err)
	assert.Nil(t, processedObject)
}
