package logger

import (
	"fmt"
	"os"
	"smicro/log/core"
	"sync"

	xerr "smicro/app/err"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger = &Logger{}

type Logger struct {
	sugar *zap.SugaredLogger
	once  sync.Once
}

func NewXLogger(c *core.Config) {
	if c == nil || c.Common == nil {
		logger.sugar = defaultXLogger()
		return
	}
	cores := make([]zapcore.Core, 0)
	cores = append(cores, zapcore.NewCore(core.DefaultConsoleEncoder(), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), core.TransformLevel(c.Common.ConsoleLevel)))
	if c.Common.IsSaveFile {
		cores = append(cores, zapcore.NewCore(core.DefaultConsoleEncoder(), core.XLogFileWriter(c.LumberJack), core.TransformLevel(c.Common.FileLevel)))
	}
	if c.Common.ConsoleStackInfo {
		core.OpenConsoleStackInfo()
	}
	log := zap.New(core.NewCores(cores...), zap.AddCaller(), zap.AddCallerSkip(1))
	logger.sugar = log.Sugar()
}

func defaultXLogger() *zap.SugaredLogger {
	xCore := zapcore.NewCore(core.DefaultConsoleEncoder(), zapcore.AddSync(os.Stdout), core.Level)
	log := zap.New(xCore, zap.AddCaller(), zap.AddCallerSkip(1))
	return log.Sugar()
}

func (x *Logger) init() {
	x.sugar = defaultXLogger()
}

func SetNamed(name string) {
	if logger.sugar == nil {
		logger.init()
	}
	logger.sugar.Named(name)
}

func SetXLoggerLevel(level core.LogLevel) {
	core.TransformLevel(level)
}

func SetXLoggerLevelString(level string) {
	core.TransformLevelString(level)
}

func Info(args ...interface{}) {
	if logger.sugar == nil {
		logger.init()
	}
	logger.sugar.Info(args...)
}

func Infof(format string, args ...interface{}) {
	if logger.sugar == nil {
		logger.init()
	}
	logger.sugar.Infof(format, args...)
}

func Warn(args ...interface{}) {
	if logger.sugar == nil {
		logger.init()
	}
	logger.sugar.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	if logger.sugar == nil {
		logger.init()
	}
	logger.sugar.Warnf(format, args...)
}

func Error(args ...interface{}) {
	if logger.sugar == nil {
		logger.init()
	}
	logger.sugar.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	if logger.sugar == nil {
		logger.init()
	}
	logger.sugar.Errorf(format, args...)
}

func XError(err *xerr.XError) {
	if logger.sugar == nil {
		logger.init()
	}
	if core.GetConsoleStack() {
		logger.sugar.Errorw(fmt.Sprintf("%+v", err.Err()), core.Stacktrace, fmt.Sprintf("%+v", err.Err()))
		return
	}
	logger.sugar.Errorw(err.Err().Error(), core.Stacktrace, fmt.Sprintf("%+v", err.Err()))
}
