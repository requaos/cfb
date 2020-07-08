package cfb

import (
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logFile := ioutil.Discard

	logger = &logrus.Logger{
		Out:       logFile,
		Formatter: new(logrus.TextFormatter),
		Level:     logrus.WarnLevel,
	}
}

func EnableDebug() {
	logger.Out = os.Stdout
	logrus.SetLevel(logrus.DebugLevel)
}

func DisableDebug() {
	logger.Out = ioutil.Discard
	logrus.SetLevel(logrus.WarnLevel)
}
