package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	ProdLevel = "prod"
)

// Logger represents a structured logger using zap.
type Logger struct {
	l *zap.Logger
}

func New() *Logger {
	lvl := os.Getenv("LOG_LEVEL")
	cfg := newZapConfig(lvl)

	logger, err := cfg.Build()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}

	return &Logger{l: logger}
}

// newZapConfig creates a zap.Config based on the given log level.
func newZapConfig(lvl string) zap.Config {
	cfg := zap.NewDevelopmentConfig()

	if lvl == ProdLevel {
		cfg = zap.NewProductionConfig()
	}

	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	return cfg
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.l.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.l.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.l.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.l.Error(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.l.Fatal(msg, fields...)
}
