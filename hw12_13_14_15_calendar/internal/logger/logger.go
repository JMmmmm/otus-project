package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const levelInfo = "INFO"
const levelDebug = "DEBUG"
const levelError = "ERROR"
const levelWarn = "WARN"

type Logger interface {
	Info(msg string)
	Debug(msg string)
	Error(msg string)
	Warn(msg string)
}

type AppLoggerAdapter struct {
	logger *zap.Logger
}

func (logger *AppLoggerAdapter) Info(msg string) {
	logger.logger.Info(msg)
}

func (logger *AppLoggerAdapter) Debug(msg string) {
	logger.logger.Debug(msg)
}

func (logger *AppLoggerAdapter) Error(msg string) {
	logger.logger.Error(msg)
}

func (logger *AppLoggerAdapter) Warn(msg string) {
	logger.logger.Warn(msg)
}

func NewAppLogger(level string, outputPath string) (Logger, error) {
	translatedLevel, err := translateLevel(level)
	if err != nil {
		return nil, err
	}
	logger, _ := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(translatedLevel),
		OutputPaths: []string{outputPath},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message", // <--
		},
	}.Build()

	return &AppLoggerAdapter{logger: logger}, nil
}

func translateLevel(level string) (zapcore.Level, error) {
	switch level {
	case levelDebug:
		return zap.DebugLevel, nil
	case levelInfo:
		return zap.InfoLevel, nil
	case levelError:
		return zap.ErrorLevel, nil
	case levelWarn:
		return zap.WarnLevel, nil
	default:
		return zap.DebugLevel, fmt.Errorf("undefined config logger level: %s", level)
	}
}
