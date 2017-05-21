package main

import "interface_logger/loggers"

// alternative methods to load the package
// import . "interface_logger/loggers"     // Import package into current namespace, no prefix required when using the methods
// import l "interface_logger/loggers"     // Import package as l, prefix methods with l

func main() {
	logger := loggers.LoggerFactory(loggers.ConsoleLoggerType)
	logger.Debug("returned from calling logger factory")
	logger.Info("writing to logger")
	logger.Warn("logger is being used a lot")
	logger.Error("program is ending soon!")
	clogger, ok := logger.(loggers.ConsoleLogger)
	if ok {
		println("The logger returned from the loggerFactory is ConsoleLogger")
		println(clogger.Name())
	}

	consoleLogger := loggers.ConsoleLogger{}
	println(consoleLogger.Name())
	consoleLogger.Info("Hello World")
}
