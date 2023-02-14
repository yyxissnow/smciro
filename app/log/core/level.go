package core

import "go.uber.org/zap/zapcore"

type XLoggerLevel int

var Level zapcore.Level

const (
	ErrorLevel XLoggerLevel = iota + 1
	WarnLevel
	InfoLevel
	DebugLevel
)

func TransformLevel(level XLoggerLevel) zapcore.Level {
	switch level {
	case ErrorLevel:
		Level = zapcore.ErrorLevel
	case WarnLevel:
		Level = zapcore.WarnLevel
	case InfoLevel:
		Level = zapcore.InfoLevel
	case DebugLevel:
		Level = zapcore.DebugLevel
	default:
		Level = zapcore.InfoLevel
	}
	return Level
}
