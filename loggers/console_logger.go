package loggers

import (
	"fmt"
)

type ConsoleLogger struct {
}

// Not exposed as Logger, would need to downcast / upcast to LinuxLogger or WindowsLoggger
func (l ConsoleLogger) Name() string {
	return "ConsoleLogger"
}

func (l ConsoleLogger) AppendMessage(mtype string, message string) {
	println(fmt.Sprintf("%s: %-40s", mtype, message))
}

func (l ConsoleLogger) Debug(message string) {
	l.AppendMessage("DEBUG", message)
}

func (l ConsoleLogger) Info(message string) {
	l.AppendMessage("INFO", message)
}

func (l ConsoleLogger) Warn(message string) {
	l.AppendMessage("WARN", message)
}

func (l ConsoleLogger) Error(message string) {
	l.AppendMessage("ERROR", message)
}
