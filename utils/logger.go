package utils

import (
	"log"
	"os"
	"sync"
	"time"
)

// Logger struct to encapsulate logging functionality
type Logger struct {
	logger *log.Logger
}

// Global logger instance
var (
	logger *Logger
	once   sync.Once
)

// InitializeLogger initializes the global logger
func InitializeLogger() {
	once.Do(func() {
		logger = &Logger{
			logger: log.New(os.Stdout, "", 0), // No default prefix
		}
	})
}

// Log formats and logs a message with timestamp, level, and section
func (l *Logger) Log(level, section, message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	l.logger.Printf("[%s] [%s] [%s] %s", timestamp, level, section, message)
}

// Convenience methods for different log levels
func (l *Logger) Info(section, message string) {
	l.Log("INFO", section, message)
}

func (l *Logger) Warn(section, message string) {
	l.Log("WARN", section, message)
}

func (l *Logger) Error(section, message string) {
	l.Log("ERROR", section, message)
}

// Expose the global logger instance
func GetLogger() *Logger {
	return logger
}
