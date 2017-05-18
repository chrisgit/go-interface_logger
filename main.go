package main

func main() {
	logger := LoggerFactory(ConsoleLoggerType)
	logger.Debug("returned from calling logger factory")
	logger.Info("writing to logger")
	logger.Warn("logger is being used a lot")
	logger.Error("program is ending soon!")
	clogger, ok := logger.(ConsoleLogger)
	if ok {
		println("The logger returned from the loggerFactory is ConsoleLogger")
		println(clogger.Name())
	}

	consoleLogger := ConsoleLogger{}
	println(consoleLogger.Name())
	consoleLogger.Info("Hello World")
}
