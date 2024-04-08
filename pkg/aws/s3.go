package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsSdkConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/codingexplorations/data-lake/pkg/log"
)

type S3Client interface {
	ListObjects(bucketName string, prefix *string) ([]types.Object, error)
	HeadObject(bucketName string, objectKey string) (*s3.HeadObjectOutput, error)
	CopyObject(srcBucket string, destBucket string, key string) (*s3.CopyObjectOutput, error)
	DeleteObjects(bucketName string, keys []string) (*s3.DeleteObjectsOutput, error)
}

type S3 struct {
	Client *s3.Client
	Logger log.Logger
}

func NewS3() (S3, error) {
	logger, err := log.GetLogger()
	if err != nil {
		return S3{}, err
	}

	cfg, err := awsSdkConfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		return S3{}, err
	}

	c := s3.NewFromConfig(cfg)

	s3Client := S3{
		Client: c,
		Logger: logger,
	}

	return s3Client, nil
}

// ListObjects lists the objects in a bucket.
func (client *S3) ListObjects(bucketName string, prefix *string) ([]types.Object, error) {
	config := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	}
	if prefix != nil {
		config.Prefix = prefix
	}

	result, err := client.Client.ListObjectsV2(context.TODO(), config)

	var contents []types.Object
	if err != nil {
		client.Logger.Error(fmt.Sprintf("couldn't list objects in bucket %v.\n", bucketName))
	} else {
		contents = result.Contents
	}

	return contents, err
}

func (client *S3) HeadObject(bucket, key string) (*s3.HeadObjectOutput, error) {
	input := &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	result, err := client.Client.HeadObject(context.TODO(), input)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (client *S3) CopyObject(srcBucket, destBucket, key string) (*s3.CopyObjectOutput, error) {
	input := &s3.CopyObjectInput{
		Bucket:     aws.String(destBucket),
		CopySource: aws.String(srcBucket + "/" + key),
		Key:        aws.String(key),
	}

	result, err := client.Client.CopyObject(context.TODO(), input)

	if err != nil {
		client.Logger.Error(fmt.Sprintf("could not copy object %v from %v to %v. Error: %v", key, srcBucket, destBucket, err))
	}

	return result, err
}

// DeleteObjects deletes a list of objects from a bucket.
func (client *S3) DeleteObjects(bucketName string, keys []string) (*s3.DeleteObjectsOutput, error) {
	var objectIds []types.ObjectIdentifier
	for _, key := range keys {
		objectIds = append(objectIds, types.ObjectIdentifier{Key: aws.String(key)})
	}
	resp, err := client.Client.DeleteObjects(context.TODO(), &s3.DeleteObjectsInput{
		Bucket: aws.String(bucketName),
		Delete: &types.Delete{Objects: objectIds},
	})
	if err != nil {
		client.Logger.Error(fmt.Sprintf("could not delete objects from bucket %v. Error: %v", bucketName, err))
	}
	return resp, err
}
