package core

import (
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"test/xlog_err/xlog"
	"time"
)

func DefaultConsoleEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:          "ts",
		LevelKey:         "level",
		NameKey:          "logger",
		CallerKey:        "caller",
		MessageKey:       "msg",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeLevel:      xLevelEncoder,
		EncodeCaller:     xCallerEncode,
		EncodeDuration:   zapcore.NanosDurationEncoder,
		EncodeTime:       xTimeEncoder,
		ConsoleSeparator: xlog.ConsoleSeparator,
	})
}

func DefaultFileEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
}

func xLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	zapcore.CapitalColorLevelEncoder(level, enc)
}

func xTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + t.Format(xlog.TimeEncoderFormat) + "]")
}

func xCallerEncode(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + caller.TrimmedPath() + "]   >>>")
}

func XLogFileWriter(c *xlog.LumberJack) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   c.Filename,
		MaxSize:    c.MaxSize,
		MaxBackups: c.MaxBackups,
		MaxAge:     c.MaxAge,
		Compress:   c.Compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}
