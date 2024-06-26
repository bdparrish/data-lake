package log

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/codingexplorations/data-lake/common/pkg/config"
	"github.com/codingexplorations/data-lake/common/pkg/models/v1/proto"
	mocks "github.com/codingexplorations/data-lake/common/test/mocks/pkg/log"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestServiceLog_Error(t *testing.T) {
	sqsClient := mocks.NewLoggerSqsClient(t)

	service := SqsLog{
		Sqs:      sqsClient,
		QueueUrl: &config.GetConfig().AwsLoggerQueueName,
	}

	output := &sqs.SendMessageOutput{
		MessageId: aws.String("00000000-0000-0000-0000-000000000001"),
	}

	// the assertion for this test is that this call is made, this test is performed by the mock object by defining this "On"
	sqsClient.On(
		"SendMessage",
		int32(0),
		map[string]types.MessageAttributeValue{},
		mock.MatchedBy(func(input string) bool {
			logResponse := proto.Log{}
			if err := protojson.Unmarshal([]byte(input), &logResponse); err != nil {
				t.Error("Error in unmarshalling log message")
			}

			return logResponse.GetMessage() == "test error" &&
				proto.Log_ERROR == logResponse.GetLevel() &&
				logResponse.GetFile() == "log/logger_sqs_test.go"
		}),
		aws.String("test-logger-queue"),
	).Return(output, nil)

	service.Error("test error")
}

func TestServiceLog_Warn(t *testing.T) {
	sqsClient := mocks.NewLoggerSqsClient(t)

	service := SqsLog{
		Sqs:      sqsClient,
		QueueUrl: &config.GetConfig().AwsLoggerQueueName,
	}

	output := &sqs.SendMessageOutput{
		MessageId: aws.String("00000000-0000-0000-0000-000000000001"),
	}

	// the assertion for this test is that this call is made, this test is performed by the mock object by defining this "On"
	sqsClient.On(
		"SendMessage",
		int32(0),
		map[string]types.MessageAttributeValue{},
		mock.MatchedBy(func(input string) bool {
			logResponse := proto.Log{}
			if err := protojson.Unmarshal([]byte(input), &logResponse); err != nil {
				t.Error("Error in unmarshalling log message")
			}

			return logResponse.GetMessage() == "test warn" &&
				proto.Log_WARNING == logResponse.GetLevel() &&
				logResponse.GetFile() == "log/logger_sqs_test.go"
		}),
		aws.String("test-logger-queue"),
	).Return(output, nil)

	service.Warn("test warn")
}

func TestServiceLog_Info(t *testing.T) {
	sqsClient := mocks.NewLoggerSqsClient(t)

	service := SqsLog{
		Sqs:      sqsClient,
		QueueUrl: &config.GetConfig().AwsLoggerQueueName,
	}

	output := &sqs.SendMessageOutput{
		MessageId: aws.String("00000000-0000-0000-0000-000000000001"),
	}

	// the assertion for this test is that this call is made, this test is performed by the mock object by defining this "On"
	sqsClient.On(
		"SendMessage",
		int32(0),
		map[string]types.MessageAttributeValue{},
		mock.MatchedBy(func(input string) bool {
			logResponse := proto.Log{}
			if err := protojson.Unmarshal([]byte(input), &logResponse); err != nil {
				t.Error("Error in unmarshalling log message")
			}

			return logResponse.GetMessage() == "test info" &&
				proto.Log_INFO == logResponse.GetLevel() &&
				logResponse.GetFile() == "log/logger_sqs_test.go"
		}),
		aws.String("test-logger-queue"),
	).Return(output, nil)

	service.Info("test info")
}

func TestServiceLog_Debug(t *testing.T) {
	sqsClient := mocks.NewLoggerSqsClient(t)

	service := SqsLog{
		Sqs:      sqsClient,
		QueueUrl: &config.GetConfig().AwsLoggerQueueName,
	}

	output := &sqs.SendMessageOutput{
		MessageId: aws.String("00000000-0000-0000-0000-000000000001"),
	}

	// the assertion for this test is that this call is made, this test is performed by the mock object by defining this "On"
	sqsClient.On(
		"SendMessage",
		int32(0),
		map[string]types.MessageAttributeValue{},
		mock.MatchedBy(func(input string) bool {
			logResponse := proto.Log{}
			if err := protojson.Unmarshal([]byte(input), &logResponse); err != nil {
				t.Error("Error in unmarshalling log message")
			}

			return logResponse.GetMessage() == "test debug" &&
				proto.Log_DEBUG == logResponse.GetLevel() &&
				logResponse.GetFile() == "log/logger_sqs_test.go"
		}),
		aws.String("test-logger-queue"),
	).Return(output, nil)

	service.Debug("test debug")
}
