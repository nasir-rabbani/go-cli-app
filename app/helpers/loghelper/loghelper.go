package loghelper

import (
	"fmt"
	"mycart/app/helpers/filehelper"
	"mycart/app/models"

	log "github.com/sirupsen/logrus"
)

var logFields = log.Fields{"client": "MyCart"}

// Init - to make one time initial setup for logrus
func Init() {
	f, err := filehelper.FilePointer(models.LogFilePath)
	if err != nil {
		fmt.Println("Can not open log file ::", err)
		return
	}
	log.SetOutput(f)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
	// log.SetOutput(os.Stdout)
}

// LogInfo logs a message at level Info.
func LogInfo(args ...interface{}) {
	log.WithFields(logFields).Info(args...)
}

// LogWarn logs a message at level Warn.
func LogWarn(args ...interface{}) {
	log.WithFields(logFields).Warn(args...)
}

// LogError logs a message at level Error.
func LogError(args ...interface{}) {
	log.WithFields(logFields).Error(args...)
}

// LogFatal logs a message at level Fatal.
func LogFatal(args ...interface{}) {
	log.WithFields(logFields).Info(args...)
}
