package lib

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func CommonLogger() *logrus.Logger {
	log := logrus.New()

	log.SetReportCaller(true)

	log.SetOutput(io.MultiWriter(os.Stdout))

	return log
}
