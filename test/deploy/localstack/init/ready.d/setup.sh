#!/bin/bash
echo "########### Setting region as env variable ##########"
export AWS_REGION=us-east-1

echo "########### Setting up localstack profile ###########"
aws configure set aws_access_key_id access_key --profile=localstack
aws configure set aws_secret_access_key secret_key --profile=localstack
aws configure set region $AWS_REGION --profile=localstack

echo "########### Setting testing profile ###########"
export AWS_DEFAULT_PROFILE=localstack

# Create buckets
awslocal s3 mb s3://$TEST_INGEST_BUCKET_NAME
awslocal s3api put-bucket-cors --bucket $TEST_INGEST_BUCKET_NAME --cors-configuration file:///etc/localstack/init/bucket-cors.json
