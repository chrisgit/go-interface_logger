package loggers

import (
	"os"
)

// Make underlying structure different
type FileLogger struct {
	FILELOC string
	LASTMSG string
}

// Not exposed as Logger, would need to downcast / upcast to LinuxLogger or WindowsLoggger
func (l FileLogger) Name() string {
	return "FileLogger"
}

func (l FileLogger) AppendMessage(mtype string, message string) {
	f, err := os.OpenFile(l.FILELOC, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(message); err != nil {
		panic(err)
	}
	l.LASTMSG = message
}

func (l FileLogger) Debug(message string) {
	l.AppendMessage("DEBUG", message)
}

func (l FileLogger) Info(message string) {
	l.AppendMessage("INFO", message)
}

func (l FileLogger) Warn(message string) {
	l.AppendMessage("WARN", message)
}

func (l FileLogger) Error(message string) {
	l.AppendMessage("ERROR", message)
}
