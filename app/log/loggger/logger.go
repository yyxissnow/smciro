package loggger

import (
	"fmt"
	xerr "github.com/yyxissnow/smicro/app/err"
	"github.com/yyxissnow/smicro/app/log/core"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type XLogger struct {
	sugar *zap.SugaredLogger
}

func New(c zapcore.Core, ops ...zap.Option) *XLogger {
	log := zap.New(c)
	log.WithOptions(ops...)
	return &XLogger{log.Sugar()}
}

func NewC(c *core.Config, ops ...zap.Option) *XLogger {
	if c == nil || c.Common == nil {
		return NewDefault(core.DefaultConsole(zapcore.InfoLevel), zapcore.ErrorLevel)
	}
	cores := make([]zapcore.Core, 0)
	cores = append(cores, zapcore.NewCore(core.DefaultConsoleEncoder(), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), c.Common.ConsoleLevel))
	if c.Common.IsSaveFile {
		cores = append(cores, zapcore.NewCore(core.DefaultConsoleEncoder(), core.XLogFileWriter(c.LumberJack), c.Common.FileLevel))
	}
	if c.Common.ConsoleStackInfo {
		core.OpenConsoleStackInfo()
	}
	log := zap.New(core.NewCores(cores...))
	log.WithOptions(ops...)
	return &XLogger{log.Sugar()}
}

func NewDefault(c zapcore.Core, lvl zapcore.Level) *XLogger {
	log := zap.New(c, zap.AddCallerSkip(core.CallerSkip), zap.AddStacktrace(lvl))
	return &XLogger{log.Sugar()}
}

func (x *XLogger) Info(args ...interface{}) {
	x.sugar.Info(args)
}

func (x *XLogger) Infof(format string, args ...interface{}) {
	x.sugar.Infof(format, args)
}

func (x *XLogger) Warn(args ...interface{}) {
	x.sugar.Warn(args)
}

func (x *XLogger) Warnf(format string, args ...interface{}) {
	x.sugar.Warnf(format, args)
}

func (x *XLogger) Error(args ...interface{}) {
	x.sugar.Error(args)
}

func (x *XLogger) Errorf(format string, args ...interface{}) {
	x.sugar.Errorf(format, args)
}

func (x *XLogger) XError(err *xerr.XError) {
	if core.GetConsoleStack() {
		x.sugar.Errorw(fmt.Sprintf("%+v", err.Err()), core.Stacktrace, fmt.Sprintf("%+v", err.Err()))
		return
	}
	x.sugar.Errorw(err.Err().Error(), core.Stacktrace, fmt.Sprintf("%+v", err.Err()))
}
