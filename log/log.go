package log

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

// Re-exporting
type Fields = logrus.Fields

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

func LoggerForPackage(module string) *logrus.Entry {
	initLogger()
	return instance.WithField("package", module)
}
