package ingest

import (
	"log"
	"strings"

	models_v1 "github.com/codingexplorations/data-lake/models/v1"
	"github.com/codingexplorations/data-lake/pkg/aws"
	"github.com/codingexplorations/data-lake/pkg/config"
)

type S3IngestProcessorImpl struct {
	conf     *config.Config
	s3Client aws.S3Client
}

func NewS3IngestProcessorImpl(conf *config.Config, s3Client aws.S3Client) *S3IngestProcessorImpl {
	return &S3IngestProcessorImpl{
		conf:     conf,
		s3Client: s3Client,
	}
}

// ProcessFolder processes the file
func (processor *S3IngestProcessorImpl) ProcessFolder(prefix string) ([]*models_v1.Object, error) {
	objects, err := processor.s3Client.ListObjects(processor.conf.AwsBucketName, &prefix)
	if err != nil {
		log.Printf("couldn't list objects in bucket %v.\n", processor.conf.AwsBucketName)
		return nil, err
	}

	processedObjects := make([]*models_v1.Object, 0)

	for _, object := range objects {
		if processedFile, err := processor.ProcessFile(*object.Key); err != nil {
			return nil, err
		} else {
			processedObjects = append(processedObjects, processedFile)
		}
	}

	return processedObjects, nil
}

// ProcessFile processes the file
func (processor *S3IngestProcessorImpl) ProcessFile(key string) (*models_v1.Object, error) {
	headObject, err := processor.s3Client.HeadObject(processor.conf.AwsBucketName, key)
	if err != nil {
		log.Printf("couldn't get object %v in bucket %v.\n", key, processor.conf.AwsBucketName)
		return nil, err
	}

	pathSplit := strings.Split(key, "/")

	object := &models_v1.Object{
		FileName:     pathSplit[len(pathSplit)-1],
		FileLocation: key,
		ContentType:  *headObject.ContentType,
		ContentSize:  int32(*headObject.ContentLength),
	}

	valid, err := validate(object)
	if err != nil {
		log.Printf("error validating object: %v\n", err)
		return nil, err
	}

	if !valid {
		log.Printf("object is invalid: %v\n", object)
		return nil, nil
	}

	return object, nil
}
