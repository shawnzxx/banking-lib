package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error
	config := zap.NewProductionConfig()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = "" // if you don't want to print stacktrace, set it to empty string
	config.EncoderConfig = encoderConfig

	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, tags ...zap.Field) {
	log.Info(message, tags...)
}

func Debug(message string, tags ...zap.Field) {
	log.Debug(message, tags...)
}

func Error(message string, tags ...zap.Field) {
	log.Error(message, tags...)
}
