package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger(app string) *logrus.Entry {
	var log = logrus.New()
	log.Formatter = new(logrus.TextFormatter)
	log.Level = logrus.TraceLevel
	log.Out = os.Stdout

	return log.WithField("app", app)
}
