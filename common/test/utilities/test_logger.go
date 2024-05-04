package utilities

import (
	"github.com/codingexplorations/data-lake/common/pkg/config"
	"golang.org/x/exp/slices"
)

type LogType int32

const (
	LogType_DEBUG LogType = 0
	LogType_INFO  LogType = 1
	LogType_WARN  LogType = 2
	LogType_ERROR LogType = 3
)

type MessageType int32

const (
	MessageType_INTERNAL MessageType = 0
	MessageType_AUDIT    MessageType = 1
)

type TestLogType struct {
	LogType     LogType
	Message     string
	MessageType MessageType
}

type TestLog struct {
	Calls []TestLogType
}

func (logger *TestLog) Error(msg string) {
	if slices.Contains([]string{"ERROR", "WARN", "INFO", "DEBUG"}, config.GetConfig().LoggerLevel) {
		logger.appendMessage(msg, LogType_ERROR, MessageType_INTERNAL)
	}
}

func (logger *TestLog) Warn(msg string) {
	if slices.Contains([]string{"WARN", "INFO", "DEBUG"}, config.GetConfig().LoggerLevel) {
		logger.appendMessage(msg, LogType_WARN, MessageType_INTERNAL)
	}
}

func (logger *TestLog) Info(msg string) {
	if slices.Contains([]string{"INFO", "DEBUG"}, config.GetConfig().LoggerLevel) {
		logger.appendMessage(msg, LogType_INFO, MessageType_INTERNAL)
	}
}

func (logger *TestLog) Debug(msg string) {
	if slices.Contains([]string{"DEBUG"}, config.GetConfig().LoggerLevel) {
		logger.appendMessage(msg, LogType_DEBUG, MessageType_INTERNAL)
	}
}

func (logger *TestLog) Audit(msg string) {
	logger.appendMessage(msg, LogType_DEBUG, MessageType_AUDIT)
}

func (logger *TestLog) appendMessage(msg string, logType LogType, msgType MessageType) {
	types := append(logger.Calls, TestLogType{
		LogType:     logType,
		Message:     msg,
		MessageType: msgType,
	})

	logger.Calls = types
}
