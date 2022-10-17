package loggers

import (
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"time"
)

// LevelHook generates different level logs file
func LevelHook(log *logrus.Logger) {
	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  "logs/info.logs",
		logrus.FatalLevel: "logs/fatal.logs",
		logrus.DebugLevel: "logs/debug.logs",
		logrus.WarnLevel:  "logs/warn.logs",
		logrus.ErrorLevel: "logs/error.logs",
		logrus.PanicLevel: "logs/panic.logs",
	}

	log.AddHook(lfshook.NewHook(
		pathMap,
		log.Formatter))
}

// RotateHook generates different level daily logs file
func RotateHook(basePath string, log *logrus.Logger) *logrus.Logger {
	infoFilePath := basePath + "/info.logs"
	infoWriter, _ := rotateLogs.New(
		infoFilePath+".%Y%m%d",
		rotateLogs.WithLinkName(infoFilePath),
		rotateLogs.WithMaxAge(7*24*time.Hour),
		rotateLogs.WithRotationTime(24*time.Hour),
	)

	fatalFilePath := basePath + "/fatal.logs"
	fatalWriter, _ := rotateLogs.New(
		fatalFilePath+".%Y%m%d",
		rotateLogs.WithLinkName(fatalFilePath),
		rotateLogs.WithMaxAge(7*24*time.Hour),
		rotateLogs.WithRotationTime(24*time.Hour),
	)

	debugFilePath := basePath + "/debug.logs"
	debugWriter, _ := rotateLogs.New(
		debugFilePath+".%Y%m%d",
		rotateLogs.WithLinkName(debugFilePath),
		rotateLogs.WithMaxAge(7*24*time.Hour),
		rotateLogs.WithRotationTime(24*time.Hour),
	)

	warningFilePath := basePath + "/warning.logs"
	warningWriter, _ := rotateLogs.New(
		warningFilePath+".%Y%m%d",
		rotateLogs.WithLinkName(warningFilePath),
		rotateLogs.WithMaxAge(7*24*time.Hour),
		rotateLogs.WithRotationTime(24*time.Hour),
	)

	errFilePath := basePath + "/err.logs"
	errWriter, _ := rotateLogs.New(
		errFilePath+".%Y%m%d",
		rotateLogs.WithLinkName(errFilePath),
		rotateLogs.WithMaxAge(7*24*time.Hour),
		rotateLogs.WithRotationTime(24*time.Hour),
	)

	panicFilePath := basePath + "/panic.logs"
	panicWriter, _ := rotateLogs.New(
		panicFilePath+".%Y%m%d",
		rotateLogs.WithLinkName(panicFilePath),
		rotateLogs.WithMaxAge(7*24*time.Hour),
		rotateLogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  infoWriter,
		logrus.FatalLevel: fatalWriter,
		logrus.DebugLevel: debugWriter,
		logrus.WarnLevel:  warningWriter,
		logrus.ErrorLevel: errWriter,
		logrus.PanicLevel: panicWriter,
	}
	log.SetReportCaller(true)
	log.AddHook(lfshook.NewHook(writeMap, &CustomFormatter{}))

	return log
}
