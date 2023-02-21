package xlog

import (
	"fmt"
	"os"
	"sync"

	xerr "github.com/yyxissnow/smicro/app/err"
	"github.com/yyxissnow/smicro/app/log/xcore"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger = &XLogger{}

type XLogger struct {
	sugar *zap.SugaredLogger
	once  sync.Once
}

func NewXLogger(c *xcore.Config) {
	if c == nil || c.Common == nil {
		logger.sugar = defaultXLogger()
		return
	}
	cores := make([]zapcore.Core, 0)
	cores = append(cores, zapcore.NewCore(xcore.DefaultConsoleEncoder(), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), xcore.TransformLevel(c.Common.ConsoleLevel)))
	if c.Common.IsSaveFile {
		cores = append(cores, zapcore.NewCore(xcore.DefaultConsoleEncoder(), xcore.XLogFileWriter(c.LumberJack), xcore.TransformLevel(c.Common.FileLevel)))
	}
	if c.Common.ConsoleStackInfo {
		xcore.OpenConsoleStackInfo()
	}
	log := zap.New(xcore.NewCores(cores...), zap.AddCaller(), zap.AddCallerSkip(1))
	logger.sugar = log.Sugar()
}

func defaultXLogger() *zap.SugaredLogger {
	xCore := zapcore.NewCore(xcore.DefaultConsoleEncoder(), zapcore.AddSync(os.Stdout), xcore.Level)
	log := zap.New(xCore, zap.AddCaller(), zap.AddCallerSkip(1))
	return log.Sugar()
}

func (x *XLogger) init() {
	x.sugar = defaultXLogger()
}

func SetNamed(name string) {
	if logger.sugar == nil {
		logger.init()
	}
	logger.sugar.Named(name)
}

func SetXLoggerLevel(level xcore.XLoggerLevel) {
	xcore.TransformLevel(level)
}

func SetXLoggerLevelString(level string) {
	xcore.TransformLevelString(level)
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
	if xcore.GetConsoleStack() {
		logger.sugar.Errorw(fmt.Sprintf("%+v", err.Err()), xcore.Stacktrace, fmt.Sprintf("%+v", err.Err()))
		return
	}
	logger.sugar.Errorw(err.Err().Error(), xcore.Stacktrace, fmt.Sprintf("%+v", err.Err()))
}
