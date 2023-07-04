package clog

var logger loggerOps = newZapLogger()

// loggerActions defines the available logging actions a logger needs to support
type loggerActions interface {
	Infof(msg string, args ...interface{})
	Warnf(msg string, args ...interface{})
	Errorf(msg string, args ...interface{})
}

// loggerOps defines the operations a logger needs to support
type loggerOps interface {
	loggerActions
	Flush()
}

// Info records an information level log
func Info(msg string) {
	logger.Infof(msg)
}

// Infof records an information level log via a formatted string
func Infof(msg string, args ...interface{}) {
	logger.Infof(msg, args...)
}

// Warn records a warning level log
func Warn(msg string) {
	logger.Warnf(msg)
}

// Warnf records a warning level log via a formatted string
func Warnf(msg string, args ...interface{}) {
	logger.Warnf(msg, args...)
}

// Error records an error level log
func Error(msg string) {
	logger.Errorf(msg)
}

// Errorf records an error level log via a formatted string
func Errorf(msg string, args ...interface{}) {
	logger.Errorf(msg, args...)
}

// Flush will ensure any in-flight logs are flushed. Should be called when the app is shutting down
func Flush() {
	logger.Flush()
}
