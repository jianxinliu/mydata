package misc

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/wailsapp/wails/v2/pkg/logger"
)

// FileLogger is a utility to log messages to a number of destinations
type CustomLogger struct {
	filename string
}

// NewFileLogger creates a new Logger.
func NewCustomLogger(filename string) logger.Logger {
	return &CustomLogger{
		filename: filename,
	}
}

func currentTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// Print works like Sprintf.
func (l *CustomLogger) Print(message string) {
	message = fmt.Sprintf("[%s] %s", currentTimeStr(), message)
	f, err := os.OpenFile(l.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.WriteString(message); err != nil {
		f.Close()
		log.Fatal(err)
	}
	f.Close()
}

func (l *CustomLogger) Println(message string) {
	l.Print(message + "\n")
}

// Trace level logging. Works like Sprintf.
func (l *CustomLogger) Trace(message string) {
	l.Println("TRACE | " + message)
}

// Debug level logging. Works like Sprintf.
func (l *CustomLogger) Debug(message string) {
	l.Println("DEBUG | " + message)
}

// Info level logging. Works like Sprintf.
func (l *CustomLogger) Info(message string) {
	l.Println("INFO  | " + message)
}

// Warning level logging. Works like Sprintf.
func (l *CustomLogger) Warning(message string) {
	l.Println("WARN  | " + message)
}

// Error level logging. Works like Sprintf.
func (l *CustomLogger) Error(message string) {
	l.Println("ERROR | " + message)
}

// Fatal level logging. Works like Sprintf.
func (l *CustomLogger) Fatal(message string) {
	l.Println("FATAL | " + message)
	os.Exit(1)
}
