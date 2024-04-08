package ingest

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/google/uuid"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/codingexplorations/data-lake/pkg/config"
	"github.com/codingexplorations/data-lake/pkg/log"
	mocks "github.com/codingexplorations/data-lake/test/mocks/pkg/aws"
	"github.com/codingexplorations/data-lake/test/utilities"
	"github.com/stretchr/testify/assert"
)

func Test_S3Processor_NewS3IngestProcessorImpl(t *testing.T) {
	conf := config.GetConfig()

	_, gormDb := utilities.NewMockDB()

	processor := NewS3IngestProcessorImpl(conf, gormDb)

	assert.NotNil(t, processor)
}

func Test_S3Processor_ProcessFolder(t *testing.T) {
	conf := config.GetConfig()

	s3Client := mocks.NewS3Client(t)

	listObjectsOutput := []types.Object{
		{
			Key: aws.String("test/test1.txt"),
		},
		{
			Key: aws.String("test/test2.txt"),
		},
	}
	s3Client.On("ListObjects", conf.AwsBucketName, aws.String("test/")).Return(listObjectsOutput, nil)

	headObjectOutput := &s3.HeadObjectOutput{
		ContentType:   aws.String("text/plain"),
		ContentLength: aws.Int64(15),
	}
	s3Client.On("HeadObject", conf.AwsBucketName, "test/test1.txt").Return(headObjectOutput, nil)
	s3Client.On("HeadObject", conf.AwsBucketName, "test/test2.txt").Return(headObjectOutput, nil)

	deleteObjectsOutput1 := &s3.DeleteObjectsOutput{
		Deleted: []types.DeletedObject{
			{
				Key: aws.String("test/test1.txt"),
			},
		},
	}
	s3Client.On("DeleteObjects", conf.AwsBucketName, []string{"test/test1.txt"}).Return(deleteObjectsOutput1, nil)
	deleteObjectsOutput2 := &s3.DeleteObjectsOutput{
		Deleted: []types.DeletedObject{
			{
				Key: aws.String("test/test2.txt"),
			},
		},
	}
	s3Client.On("DeleteObjects", conf.AwsBucketName, []string{"test/test2.txt"}).Return(deleteObjectsOutput2, nil)

	mockDb, gormDb := utilities.NewMockDB()
	objectsRow1 := sqlmock.NewRows(
		[]string{"id", "file_name", "file_location", "content_type", "content_size"},
	).AddRow(
		uuid.New().String(), "test1.txt", "test/test1.txt", "text/plain", 15,
	)
	objectsRow2 := sqlmock.NewRows(
		[]string{"id", "file_name", "file_location", "content_type", "content_size"},
	).AddRow(
		uuid.New().String(), "test2.txt", "test/test2.txt", "text/plain", 15,
	)
	mockDb.ExpectQuery("^INSERT INTO \"objects\".*").WithArgs("test1.txt", "test/test1.txt", "text/plain", 15).WillReturnRows(objectsRow1)
	mockDb.ExpectQuery("^INSERT INTO \"objects\".*").WithArgs("test2.txt", "test/test2.txt", "text/plain", 15).WillReturnRows(objectsRow2)

	processor := &S3IngestProcessorImpl{
		conf:     conf,
		logger:   log.NewConsoleLog(),
		s3Client: s3Client,
		db:       gormDb,
	}

	processedObjects, err := processor.ProcessFolder("test/")

	assert.Nil(t, err)
	assert.Equal(t, 2, len(processedObjects))
	assert.Equal(t, "test1.txt", processedObjects[0].FileName)
	assert.Equal(t, "test/test1.txt", processedObjects[0].FileLocation)
	assert.Equal(t, "text/plain", processedObjects[0].ContentType)
	assert.Equal(t, int32(15), processedObjects[0].ContentSize)
	assert.Equal(t, "test2.txt", processedObjects[1].FileName)
	assert.Equal(t, "test/test2.txt", processedObjects[1].FileLocation)
	assert.Equal(t, "text/plain", processedObjects[1].ContentType)
	assert.Equal(t, int32(15), processedObjects[1].ContentSize)
}

func Test_S3Processor_ProcessFile(t *testing.T) {
	conf := config.GetConfig()

	s3Client := mocks.NewS3Client(t)

	headObjectOutput := &s3.HeadObjectOutput{
		ContentType:   aws.String("text/plain"),
		ContentLength: aws.Int64(15),
	}
	s3Client.On("HeadObject", conf.AwsBucketName, "test/test.txt").Return(headObjectOutput, nil)

	deleteObjectsOutput1 := &s3.DeleteObjectsOutput{
		Deleted: []types.DeletedObject{
			{
				Key: aws.String("test/test.txt"),
			},
		},
	}
	s3Client.On("DeleteObjects", conf.AwsBucketName, []string{"test/test.txt"}).Return(deleteObjectsOutput1, nil)

	mockDb, gormDb := utilities.NewMockDB()
	objectsRow := sqlmock.NewRows(
		[]string{"id", "file_name", "file_location", "content_type", "content_size"},
	).AddRow(
		uuid.New().String(), "test.txt", "test/test.txt", "text/plain", 15,
	)
	mockDb.ExpectQuery("^INSERT INTO \"objects\".*").WithArgs().WillReturnRows(objectsRow)

	processor := &S3IngestProcessorImpl{
		conf:     conf,
		logger:   log.NewConsoleLog(),
		s3Client: s3Client,
		db:       gormDb,
	}

	processedObject, err := processor.ProcessFile("test/test.txt")

	assert.Nil(t, err)
	assert.Equal(t, "test.txt", processedObject.FileName)
	assert.Equal(t, "test/test.txt", processedObject.FileLocation)
	assert.Equal(t, "text/plain", processedObject.ContentType)
	assert.Equal(t, int32(15), processedObject.ContentSize)
}
