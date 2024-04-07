package ingest

import (
	"os"
	"path/filepath"
	"testing"

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
			processor := &LocalIngestProcessorImpl{}

			pwd, _ := os.Getwd()
			_ = os.MkdirAll(pwd+tc.folder+"/"+tc.subfolder, os.ModePerm)

			d1 := []byte("This is a test.")
			fileName := filepath.Join(pwd, tc.folder, tc.subfolder, "test.txt")
			_ = os.WriteFile(fileName, d1, 0644)

			processedObjects, err := processor.ProcessFolder(pwd + tc.folder)

			assert.Nil(t, err)
			assert.Len(t, processedObjects, 1)
			assert.Equal(t, "test.txt", processedObjects[0].FileName)
			assert.Equal(t, filepath.Join(pwd, tc.folder, tc.subfolder, tc.location), processedObjects[0].FileLocation)
			assert.Equal(t, "text/plain", processedObjects[0].ContentType)
			assert.Equal(t, int32(15), processedObjects[0].ContentSize)
		})
	}
}

func TestFolderIngest_ProcessFile_Success(t *testing.T) {
	processor := &LocalIngestProcessorImpl{}

	pwd, _ := os.Getwd()
	_ = os.Mkdir("test", os.ModePerm)

	d1 := []byte("test go")
	fileName := filepath.Join(pwd, "test", "test.txt")
	_ = os.WriteFile(fileName, d1, 0644)

	processedObject, err := processor.ProcessFile(fileName)

	assert.Nil(t, err)
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
