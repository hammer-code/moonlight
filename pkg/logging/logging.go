package logging

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"
)

type Fields map[string]interface{}

func InitLogging() {
	// Log as JSON instead of the default ASCII formatter.
	// log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func Info(ctx context.Context, message string, fields ...Fields) {
	newLogfileds := mergeField(fields...)
	log.WithContext(ctx).WithFields(newLogfileds).Info(message)
}

func Error(ctx context.Context, err error, message string, fields ...Fields) {
	newLogfileds := mergeField(fields...)
	log.WithContext(ctx).WithFields(newLogfileds).WithError(err).Error(message)
}

func Debug(ctx context.Context, message string, fields ...Fields) {
	newLogfileds := mergeField(fields...)
	log.WithContext(ctx).WithFields(newLogfileds).Debug(message)
}

func mergeField(logFields ...Fields) log.Fields {
	newLogfields := make(log.Fields)

	for _, field := range logFields {
		for key, value := range field {
			newLogfields[key] = value
		}
	}

	return newLogfields
}
