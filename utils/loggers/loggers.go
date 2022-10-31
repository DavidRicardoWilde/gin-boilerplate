package loggers

import (
	"bytes"
	"fmt"
	"gin-boilerplate/utils/configs"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"time"
)

var ApiLog *logrus.Logger
var ScheduleLog *logrus.Logger
var logPath string

func init() {
	logPath = "logs/"
}

// InitApiServerLog initializes the log group for api server
func InitApiServerLog() {
	groupName := "api"
	ApiLog = logInit(groupName)
}

// InitScheduleLog initializes the log group for schedule
func InitScheduleLog() {
	groupName := "schedule"
	ScheduleLog = logInit(groupName)
}

type CustomFormatter struct{}

// Format sets custom format for logrus
func (m *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	entry.Time = entry.Time.In(time.UTC) // It is recommended to use UTC time zone instead of Local, it is also possible to use the specified time zone
	// example of using other time zone: PST, CST and Local
	//entry.Time = entry.Time.In(time.FixedZone("PST", -8*60*60))
	//entry.Time = entry.Time.In(time.FixedZone("CST", 8*60*60))
	//entry.Time = entry.Time.In(time.Local)
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	var newLog string
	if entry.Caller != nil {
		fName := filepath.Base(entry.Caller.File)
		newLog = fmt.Sprintf("%s [%s] [%s:%d %s] %s\n", timestamp, entry.Level.String(), fName, entry.Caller.Line, entry.Caller.Function, entry.Message)
	} else {
		newLog = fmt.Sprintf("%s [%s] %s", timestamp, entry.Level.String(), entry.Message)
	}

	b.WriteString(newLog)
	return b.Bytes(), nil
}

// logInit initializes and sets log config including log level, log files path and hooks.
func logInit(groupName string) *logrus.Logger {
	log := logrus.New()
	log = RotateHook(logPath+groupName, log)
	log.SetLevel(getLogLevel()) // Set the logs level which will be output.

	log.Infof("Log level: %s", configs.GetGlobalAppConfig().LogLevel)

	return log
}

// getLogLevel returns the log level according to the configuration file
func getLogLevel() logrus.Level {
	switch configs.GetGlobalAppConfig().LogLevel {
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	default:
		return logrus.InfoLevel
	}
}
