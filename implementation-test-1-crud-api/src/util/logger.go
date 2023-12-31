package util

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.Out = os.Stdout
	log.SetFormatter(&logrus.JSONFormatter{})
}

func Error(err error, message interface{}) {
	if message != nil {
		log.WithFields(logrus.Fields{
			"message": message,
		})
	}
	log.Error(err)
}

func Info(info, message interface{}) {
	if message != nil {
		log.WithFields(logrus.Fields{
			"message": message,
		})
	}
	log.Info(info)
}
