package pkg

import (
	"os"
	"testing"

	modelsv1 "github.com/codeexplorations/data-lake/models/v1"
	"github.com/stretchr/testify/assert"
)

func TestFolderIngest_ProcessFolder_CheckDepth(t *testing.T) {
	tests := []struct {
		name     string
		folder   string
		location string
	}{
		{
			name:     "directory",
			folder:   "/../test",
			location: "/files/test.txt",
		},
		{
			name:     "file",
			folder:   "/../test/files",
			location: "/test.txt",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pwd, _ := os.Getwd()

			processedObjects, err := ProcessFolder(pwd + tc.folder)

			assert.Nil(t, err)
			assert.Len(t, processedObjects, 1)
			assert.Equal(t, "test.txt", processedObjects[0].FileName)
			assert.Equal(t, pwd+tc.folder+tc.location, processedObjects[0].FileLocation)
			assert.Equal(t, "text/plain", processedObjects[0].ContentType)
			assert.Equal(t, int32(15), processedObjects[0].ContentSize)
		})
	}
}

func TestFolderIngest_ProcessFile_Success(t *testing.T) {
	pwd, _ := os.Getwd()

	fileName := pwd + "/../test/files/test.txt"

	processedObject, err := ProcessFile(fileName)

	assert.Nil(t, err)
	assert.Equal(t, "test.txt", processedObject.FileName)
	assert.Equal(t, fileName, processedObject.FileLocation)
	assert.Equal(t, "text/plain", processedObject.ContentType)
	assert.Equal(t, int32(15), processedObject.ContentSize)
}

func TestFolderIngest_ProcessFolder_Failure(t *testing.T) {
	pwd, _ := os.Getwd()

	folder := pwd + "/../../test/files/missing"

	processedObject, err := ProcessFolder(folder)

	assert.Error(t, err)
	assert.Equal(t, "open /Users/benjaminparrish/Development/Personal/data-lake/pkg/../../test/files/missing: no such file or directory", err.Error())
	assert.Nil(t, processedObject)
}

func TestFolderIngest_ProcessFile_Failure(t *testing.T) {
	pwd, _ := os.Getwd()

	fileName := pwd + "/../../test/files/missing.txt"

	processedObject, err := ProcessFile(fileName)

	assert.Error(t, err)
	assert.Nil(t, processedObject)
}

func TestFolderIngest_ProcessFile_validate(t *testing.T) {
	object := &modelsv1.Object{
		FileName:     "test.txt",
		FileLocation: "/tmp/test/test.txt",
		ContentType:  "text/plain",
		ContentSize:  15,
	}

	valid, err := validate(object)

	assert.Nil(t, err)
	assert.True(t, valid)
}

func TestFolderIngest_ProcessFile_validateTable(t *testing.T) {
	tests := []struct {
		name          string
		object        *modelsv1.Object
		expectedError string
	}{
		{
			name: "invalid - FileName empty",
			object: &modelsv1.Object{
				FileName:     "",
				FileLocation: "/tmp/test/test.txt",
				ContentType:  "text/plain",
				ContentSize:  15,
			},
			expectedError: "file_name: value is required [required]",
		},
		{
			name: "invalid - FileLocation empty",
			object: &modelsv1.Object{
				FileName:     "test.txt",
				FileLocation: "",
				ContentType:  "text/plain",
				ContentSize:  15,
			},
			expectedError: "file_location: value is required [required]",
		},
		{
			name: "invalid - ContentType empty",
			object: &modelsv1.Object{
				FileName:     "test.txt",
				FileLocation: "/tmp/test/test.txt",
				ContentType:  "",
				ContentSize:  15,
			},
			expectedError: "content_type: value is required [required]",
		},
		{
			name: "invalid - ContentSize less than 0",
			object: &modelsv1.Object{
				FileName:     "test.txt",
				FileLocation: "/tmp/test/test.txt",
				ContentType:  "text/plain",
				ContentSize:  0,
			},
			expectedError: "content_size: value must be greater than 0 and less than or equal to 1048576 [int32.gt_lte]",
		},
		{
			name: "invalid - ContentSize greater than 1GB",
			object: &modelsv1.Object{
				FileName:     "test.txt",
				FileLocation: "/tmp/test/test.txt",
				ContentType:  "text/plain",
				ContentSize:  2097152,
			},
			expectedError: "content_size: value must be greater than 0 and less than or equal to 1048576 [int32.gt_lte]",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			valid, err := validate(tc.object)

			assert.Error(t, err)
			assert.False(t, valid)
			assert.Contains(t, err.Error(), tc.expectedError)
		})
	}
}
