package core

import (
	"go.uber.org/zap/zapcore"
)

var consoleStackInfo = false //终端是否打印堆栈信息

type Config struct {
	Common     *Common
	LumberJack *LumberJack
}

type Common struct {
	ConsoleLevel     zapcore.Level
	FileLevel        zapcore.Level
	StackLevel       zapcore.Level
	ConsoleStackInfo bool
	IsSaveFile       bool
}

type LumberJack struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

func OpenConsoleStackInfo() {
	consoleStackInfo = true
}

func CloseConsoleStackInfo() {
	consoleStackInfo = false
}

func GetConsoleStack() bool {
	return consoleStackInfo
}

func ChangeTimeEncoderFormat(format string) {
	TimeEncoderFormat = format
}

func DefaultConfig() *Config {
	return &Config{
		Common: &Common{
			ConsoleLevel:     zapcore.InfoLevel,
			FileLevel:        zapcore.InfoLevel,
			StackLevel:       zapcore.ErrorLevel,
			ConsoleStackInfo: false,
			IsSaveFile:       false,
		}}
}
