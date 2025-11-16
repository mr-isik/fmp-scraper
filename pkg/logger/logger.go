package logger

import (
	"log"
	"os"
)

// Logger defines the interface for logging operations
type Logger interface {
	Info(msg string)
	Error(msg string)
	Success(msg string)
	Warning(msg string)
}

// ConsoleLogger implements console-based logging
type ConsoleLogger struct {
	infoLogger    *log.Logger
	errorLogger   *log.Logger
	successLogger *log.Logger
	warningLogger *log.Logger
}

// NewConsoleLogger creates a new console logger
func NewConsoleLogger() *ConsoleLogger {
	return &ConsoleLogger{
		infoLogger:    log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime),
		errorLogger:   log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime),
		successLogger: log.New(os.Stdout, "[SUCCESS] ", log.Ldate|log.Ltime),
		warningLogger: log.New(os.Stdout, "[WARNING] ", log.Ldate|log.Ltime),
	}
}

// Info logs an informational message
func (l *ConsoleLogger) Info(msg string) {
	l.infoLogger.Println(msg)
}

// Error logs an error message
func (l *ConsoleLogger) Error(msg string) {
	l.errorLogger.Println(msg)
}

// Success logs a success message
func (l *ConsoleLogger) Success(msg string) {
	l.successLogger.Println(msg)
}

// Warning logs a warning message
func (l *ConsoleLogger) Warning(msg string) {
	l.warningLogger.Println(msg)
}

// Errorf logs a formatted error message
func (l *ConsoleLogger) Errorf(format string, args ...interface{}) {
	l.errorLogger.Printf(format, args...)
}

// Infof logs a formatted informational message
func (l *ConsoleLogger) Infof(format string, args ...interface{}) {
	l.infoLogger.Printf(format, args...)
}
