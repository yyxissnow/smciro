package core

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func NewCore(enc zapcore.Encoder, ws zapcore.WriteSyncer, lev zapcore.LevelEnabler) zapcore.Core {
	return wrapXCore(zapcore.NewCore(enc, ws, lev))
}

func NewCores(cores ...zapcore.Core) zapcore.Core {
	for i := range cores {
		cores[i] = wrapXCore(cores[i])
	}
	return zapcore.NewTee(cores...)
}

func DefaultConsole(l zapcore.Level) zapcore.Core {
	return wrapXCore(zapcore.NewCore(
		DefaultConsoleEncoder(),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
		zap.NewAtomicLevelAt(l)),
	)
}

func DefaultFile(c *LumberJack, l zapcore.Level) zapcore.Core {
	return wrapXCore(zapcore.NewCore(
		DefaultFileEncoder(),
		zapcore.NewMultiWriteSyncer(zapcore.NewMultiWriteSyncer(XLogFileWriter(c))),
		zap.NewAtomicLevelAt(l)),
	)
}

type xCore struct { //封装新的日志核心
	zapcore.Core
}

func wrapXCore(c zapcore.Core) zapcore.Core {
	return &xCore{c}
}

func (x *xCore) With(fields []zapcore.Field) zapcore.Core {
	return &xCore{x.Core.With(fields)}
}

func (x *xCore) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	if !hasStackErr(fields) { //写入之前进行判断额外字段中是否包含stacktrace字段
		return x.Core.Write(ent, fields) //没有不做处理
	}
	ent.Stack, fields = getStacks(fields) //有则将字段抽出来放入默认字段中
	return x.Core.Write(ent, fields)
}

func (x *xCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if x.Enabled(ent.Level) {
		return ce.AddCore(ent, x)
	}
	return ce
}

func (x *xCore) Sync() error {
	return x.Core.Sync()
}

func hasStackErr(fields []zapcore.Field) bool {
	for _, field := range fields {
		if field.Key == Stacktrace {
			return true
		}
	}
	return false
}

func getStacks(fields []zapcore.Field) (string, []zapcore.Field) {
	for i, field := range fields {
		if field.Key == Stacktrace {
			fields = append(fields[:i], fields[i+1:]...)
			return field.String, fields
		}
	}
	return "", fields
}
