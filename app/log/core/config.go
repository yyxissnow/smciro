package core

var consoleStackInfo = false //终端是否打印堆栈信息

type Config struct {
	Common     *Common
	LumberJack *LumberJack
}

type Common struct {
	ConsoleLevel     XLoggerLevel
	FileLevel        XLoggerLevel
	StackLevel       XLoggerLevel
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
			ConsoleLevel:     InfoLevel,
			FileLevel:        InfoLevel,
			StackLevel:       ErrorLevel,
			ConsoleStackInfo: false,
			IsSaveFile:       false,
		}}
}
