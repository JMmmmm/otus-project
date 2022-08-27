package logger

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

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
	lvl, err := zap.ParseAtomicLevel(strings.ToLower(level))
	if err != nil {
		return nil, fmt.Errorf("can't initialize level zap logger: %w", err)
	}
	logger, _ := zap.Config{
		Encoding:    "json",
		Level:       lvl,
		OutputPaths: []string{outputPath},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",
			LevelKey:   "level-key",
		},
	}.Build()

	return &AppLoggerAdapter{logger: logger}, nil
}
