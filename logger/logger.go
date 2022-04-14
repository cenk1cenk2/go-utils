package logger

import (
	"os"

	"github.com/sirupsen/logrus"
	formatter "gitlab.kilic.dev/libraries/go-utils/logger/formatter"
)

// Log Returns a new logrus logger instance.
var Log = logrus.New()

// InitiateLogger the default logger
func InitiateLogger(level logrus.Level) *logrus.Logger {
	Log.Out = os.Stdout

	Log.SetFormatter(&formatter.Formatter{
		FieldsOrder:      []string{"context"},
		TimestampFormat:  "",
		HideKeys:         true,
		NoColors:         false,
		NoFieldsColors:   false,
		NoFieldsSpace:    false,
		ShowFullLevel:    false,
		NoUppercaseLevel: false,
		TrimMessages:     true,
		CallerFirst:      false,
	})

	Log.SetLevel(level)

	return Log
}
