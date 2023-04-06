package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = NewLogger()

func NewLogger() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.ErrorLevel)

	return log
}
