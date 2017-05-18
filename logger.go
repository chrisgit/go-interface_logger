package main

// Greeting messages
const (
	ConsoleLoggerType = "Console"
	FileLoggerType    = "File"
)

type Logger interface {
	Debug(string)
	Info(string)
	Warn(string)
	Error(string)
}

func LoggerFactory(logType string) Logger {
	if logType == FileLoggerType {
		return FileLogger{"C:\\Temp\\log.txt", ""}
	}
	return ConsoleLogger{}
}
