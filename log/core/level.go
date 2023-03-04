package core

import "go.uber.org/zap/zapcore"

type LogLevel int

var Level zapcore.Level

const (
	ErrorLevel LogLevel = iota + 1
	WarnLevel
	InfoLevel
	DebugLevel

	ErrorLevelString = "error"
	WarnLevelString  = "warn"
	InfoLevelString  = "info"
	DebugLevelSting  = "debug"
)

func TransformLevel(level LogLevel) zapcore.Level {
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

func TransformLevelString(level string) zapcore.Level {
	switch level {
	case ErrorLevelString:
		Level = zapcore.ErrorLevel
	case WarnLevelString:
		Level = zapcore.WarnLevel
	case InfoLevelString:
		Level = zapcore.InfoLevel
	case DebugLevelSting:
		Level = zapcore.DebugLevel
	default:
		Level = zapcore.InfoLevel
	}
	return Level
}
