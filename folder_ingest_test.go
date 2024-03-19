package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFolderIngest_ProcessFolder_SuccessNoDepth(t *testing.T) {
	pwd, _ := os.Getwd()

	folder := pwd + "/test/files"

	processedObjects, err := ProcessFolder(folder)

	assert.Nil(t, err)
	assert.Len(t, processedObjects, 1)
	assert.Equal(t, "test.txt", processedObjects[0].FileName)
	assert.Equal(t, folder+"/test.txt", processedObjects[0].FileLocation)
	assert.Equal(t, "text/plain", processedObjects[0].ContentType)
	assert.Equal(t, int32(15), processedObjects[0].ContentSize)
}

func TestFolderIngest_ProcessFolder_SuccessWithDepth(t *testing.T) {
	pwd, _ := os.Getwd()

	folder := pwd + "/test"

	processedObjects, err := ProcessFolder(folder)

	assert.Nil(t, err)
	assert.Len(t, processedObjects, 1)
	assert.Equal(t, "test.txt", processedObjects[0].FileName)
	assert.Equal(t, folder+"/files/test.txt", processedObjects[0].FileLocation)
	assert.Equal(t, "text/plain", processedObjects[0].ContentType)
	assert.Equal(t, int32(15), processedObjects[0].ContentSize)
}

func TestFolderIngest_ProcessFile_Success(t *testing.T) {
	pwd, _ := os.Getwd()

	fileName := pwd + "/test/files/test.txt"

	processedObject, err := ProcessFile(fileName)

	assert.Nil(t, err)
	assert.Equal(t, "test.txt", processedObject.FileName)
	assert.Equal(t, fileName, processedObject.FileLocation)
	assert.Equal(t, "text/plain", processedObject.ContentType)
	assert.Equal(t, int32(15), processedObject.ContentSize)
}
