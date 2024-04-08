package ingest

import (
	"fmt"
	golog "log"
	"strings"

	dbModels "github.com/codingexplorations/data-lake/models/v1/db"
	"github.com/codingexplorations/data-lake/pkg/aws"
	"github.com/codingexplorations/data-lake/pkg/config"
	"github.com/codingexplorations/data-lake/pkg/log"
	"gorm.io/gorm"
)

type S3IngestProcessorImpl struct {
	conf     *config.Config
	logger   log.Logger
	s3Client aws.S3Client
	db       *gorm.DB
}

func NewS3IngestProcessorImpl(conf *config.Config, db *gorm.DB) *S3IngestProcessorImpl {
	logger, err := log.GetLogger()
	if err != nil {
		golog.Fatalf("couldn't create logger: %v\n", err)
	}

	logger.Info("Using S3 ingest processor")

	s3Client, err := aws.NewS3()
	if err != nil {
		logger.Error(fmt.Sprintf("couldn't create s3 client: %v\n", err))
		return nil
	}

	return &S3IngestProcessorImpl{
		conf:     conf,
		logger:   logger,
		s3Client: &s3Client,
		db:       db,
	}
}

// ProcessFolder processes the file
func (processor *S3IngestProcessorImpl) ProcessFolder(prefix string) ([]*dbModels.Object, error) {
	processor.logger.Info(fmt.Sprintf("Processing folder: %s", prefix))

	objects, err := processor.s3Client.ListObjects(processor.conf.AwsIngestBucketName, &prefix)
	if err != nil {
		processor.logger.Error(fmt.Sprintf("couldn't list objects in bucket %v.\n", processor.conf.AwsIngestBucketName))
		return nil, err
	}

	processedObjects := make([]*dbModels.Object, 0)

	for _, object := range objects {
		if processedFile, err := processor.ProcessFile(*object.Key); err != nil {
			return nil, err
		} else {
			processor.logger.Info(fmt.Sprintf("processed file: %v\n", processedFile))
			processedObjects = append(processedObjects, processedFile)
		}
	}

	return processedObjects, nil
}

// ProcessFile processes the file
func (processor *S3IngestProcessorImpl) ProcessFile(key string) (*dbModels.Object, error) {
	processor.logger.Info(fmt.Sprintf("Processing key: %s", key))

	headObject, err := processor.s3Client.HeadObject(processor.conf.AwsIngestBucketName, key)
	if err != nil {
		processor.logger.Error(fmt.Sprintf("couldn't get object %v in bucket %v.\n", key, processor.conf.AwsIngestBucketName))
		return nil, err
	}

	pathSplit := strings.Split(key, "/")

	object := &dbModels.Object{
		FileName:     pathSplit[len(pathSplit)-1],
		FileLocation: key,
		ContentType:  *headObject.ContentType,
		ContentSize:  int32(*headObject.ContentLength),
	}

	valid, err := validate(object)
	if err != nil {
		processor.logger.Error(fmt.Sprintf("error validating object: %v\n", err))
		return nil, err
	}

	if !valid {
		processor.logger.Error(fmt.Sprintf("object is invalid: %v\n", object))
		return nil, nil
	}

	_, err = processor.s3Client.CopyObject(processor.conf.AwsIngestBucketName, processor.conf.AwsCatalogBucketName, key)
	if err != nil {
		processor.logger.Error(fmt.Sprintf("couldn't copy object %v from bucket %v to bucket %v.\n", key, processor.conf.AwsIngestBucketName, processor.conf.AwsCatalogBucketName))
		return nil, err
	}

	deleteObjectsOutput, err := processor.s3Client.DeleteObjects(processor.conf.AwsIngestBucketName, []string{key})
	if err != nil {
		processor.logger.Error(fmt.Sprintf("couldn't delete object %v in bucket %v.\n", key, processor.conf.AwsIngestBucketName))
		return nil, err
	}

	if len(deleteObjectsOutput.Deleted) == 0 {
		processor.logger.Error(fmt.Sprintf("couldn't delete object %v in bucket %v.\n", key, processor.conf.AwsIngestBucketName))
		return nil, nil
	}

	if err := processor.db.Create(&object).Error; err != nil {
		processor.logger.Error(fmt.Sprintf("error creating object record - %v", err))
		return nil, err
	}

	return object, nil
}
