package ingest

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/codingexplorations/data-lake/pkg/config"
	"github.com/codingexplorations/data-lake/pkg/log"
	mocks "github.com/codingexplorations/data-lake/test/mocks/pkg/aws"
	"github.com/stretchr/testify/assert"
)

func Test_S3Processor_ProcessFolder(t *testing.T) {
	conf := config.GetConfig()

	s3Client := mocks.NewS3Client(t)

	headObjectOutput := &s3.HeadObjectOutput{}
	s3Client.On("HeadObject", conf.AwsBucketName, "test/test.txt").Return(headObjectOutput, nil)

	processor := &S3IngestProcessorImpl{
		conf:     conf,
		logger:   log.NewConsoleLog(),
		s3Client: s3Client,
	}
	processor.s3Client = s3Client

	processedObject, err := processor.ProcessFile("test/test.txt")

	assert.Nil(t, err)
	assert.Equal(t, "test.txt", processedObject.FileName)
	assert.Equal(t, "test/test.txt", processedObject.FileLocation)
	assert.Equal(t, "text/plain", processedObject.ContentType)
	assert.Equal(t, int32(15), processedObject.ContentSize)
}
