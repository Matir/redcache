package log

import (
	"io"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

// Re-exporting
type Fields = logrus.Fields
type Level = logrus.Level

var (
	instance *logrus.Logger
	initOnce sync.Once
)

func initLogger() {
	initOnce.Do(func() {
		instance = &logrus.Logger{
			Out:       os.Stderr,
			Formatter: new(logrus.TextFormatter),
			Hooks:     make(logrus.LevelHooks),
			Level:     logrus.DebugLevel,
		}
	})
}

func LoggerForPackage(pkg string) *logrus.Entry {
	initLogger()
	return instance.WithField("package", pkg)
}

func SetOutput(w io.Writer) {
	instance.Out = w
}

func SetLevel(l Level) {
	instance.Level = l
}

func ParseLevel(l string) (Level, error) {
	return logrus.ParseLevel(l)
}
