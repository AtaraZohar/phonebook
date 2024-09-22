package logger

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func Init() {
	logFile, err := os.OpenFile("/app/logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	Log = logrus.New()
	Log.Out = logFile

	Log.Info("Logger initialized")
}
