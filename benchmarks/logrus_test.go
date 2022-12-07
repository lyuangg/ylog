package benchmarks

import (
	"io"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

func newLogrusDefaultLogger() *logrus.Logger {
	return &logrus.Logger{
		Out:       io.Discard,
		Formatter: new(logrus.TextFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}
}

func newLogrusRoateLogger() *logrus.Logger {
	l := logrus.New()

	path := "./testdata/logrus/testlog"
	days := 1
	writer, _ := rotatelogs.New(
		path+"-%Y-%m-%d"+".log",
		rotatelogs.WithLinkName(path+".log"),
		rotatelogs.WithMaxAge(-1),
		rotatelogs.WithRotationCount(uint(days)),
		rotatelogs.WithRotationTime(time.Hour*time.Duration(24)),
	)
	l.SetFormatter(&logrus.TextFormatter{})
	l.SetOutput(writer)

	return l
}

func newLogrusJsonLogger() *logrus.Logger {
	return &logrus.Logger{
		Out:       io.Discard,
		Formatter: new(logrus.JSONFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}
}
