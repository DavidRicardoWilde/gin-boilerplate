package loggers

import (
	"bytes"
	"fmt"
	"gin-boilerplate/utils/configs"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"time"
)

var Log *logrus.Logger
var ScheduleLog *logrus.Logger

func InitWebServerLog() {
	Log = logInit("logs/web")
}

func InitScheduleLog() {
	ScheduleLog = logInit("logs/schedule")
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

	entry.Time = entry.Time.In(time.UTC)
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

func logInit(group string) *logrus.Logger {
	log := logrus.New()
	log = RotateHook(group, log)
	log.SetLevel(getLogLevel()) // Set the logs level which will be output.

	log.Infof("Log level: %s", configs.GetGlobalAppConfig().LogLevel)

	return log
}

func getLogLevel() logrus.Level {
	if configs.GetGlobalAppConfig().LogLevel == "dev" {
		return logrus.DebugLevel
	}

	if configs.GetGlobalAppConfig().LogLevel == "prod" {
		return logrus.DebugLevel
	}

	return logrus.InfoLevel
}
