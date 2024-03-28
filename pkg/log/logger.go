package log

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/codingexplorations/data-lake/pkg/config"
	"golang.org/x/exp/slices"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Yellow = "\033[33m"
var Cyan = "\033[36m"

type Logger interface {
	Error(msg string)
	Warn(msg string)
	Info(msg string)
	Debug(msg string)
}

type ConsoleLog struct{}

func NewConsoleLog() *ConsoleLog {
	return &ConsoleLog{}
}

func (logger *ConsoleLog) Error(msg string) {
	if slices.Contains([]string{"ERROR", "WARN", "INFO", "DEBUG"}, config.GetConfig().LoggingLevel) {
		file, line := getCaller(2)
		msg = fmt.Sprintf("[ERROR] %s#%d - %s", file, line, msg)
		log.Println(Red + msg + Reset)
	}
}

func (logger *ConsoleLog) Warn(msg string) {
	if slices.Contains([]string{"WARN", "INFO", "DEBUG"}, config.GetConfig().LoggingLevel) {
		file, line := getCaller(2)
		msg = fmt.Sprintf("[WARN] %s#%d - %s", file, line, msg)
		log.Println(Yellow + msg + Reset)
	}
}

func (logger *ConsoleLog) Info(msg string) {
	if slices.Contains([]string{"INFO", "DEBUG"}, config.GetConfig().LoggingLevel) {
		file, line := getCaller(2)
		msg = fmt.Sprintf("[INFO] %s#%d - %s", file, line, msg)
		log.Println(msg)
	}
}

func (logger *ConsoleLog) Debug(msg string) {
	if slices.Contains([]string{"DEBUG"}, config.GetConfig().LoggingLevel) {
		file, line := getCaller(2)
		msg = fmt.Sprintf("[DEBUG] %s#%d - %s", file, line, msg)
		log.Println(Cyan + msg + Reset)
	}
}

func getCaller(skip int) (string, int32) {
	_, file, line, ok := runtime.Caller(skip)

	if ok {
		callerSplit := strings.Split(file, "/")

		// get the last two elements in the file path on the / split - Go splice range functionality
		lastTwoFilePaths := callerSplit[len(callerSplit)-2:]

		shortFile := strings.Join(lastTwoFilePaths, "/")

		return shortFile, int32(line)
	}

	return "", int32(line)
}
