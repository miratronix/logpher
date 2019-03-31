package logpher

import (
	"fmt"
	"strings"
	"time"
)

// logger Defines a logger structure
type Logger struct {
	Logpher *Logpher `autumn:"logpher"`
	name    string
	level   *level
}

// newLogger constructs a logger with the specified name, level, and writer
func newLogger(name string, logpher *Logpher) *Logger {
	l := &Logger{Logpher: logpher, name: name}
	l.PostConstruct()
	return l
}

// Trace logs at the trace level
func (l *Logger) Trace(data ...interface{}) {
	l.log(traceLevel, data...)
}

// Debug logs at the debug level
func (l *Logger) Debug(data ...interface{}) {
	l.log(debugLevel, data...)
}

// Info logs at the info level
func (l *Logger) Info(data ...interface{}) {
	l.log(infoLevel, data...)
}

// Warn logs at the warn level
func (l *Logger) Warn(data ...interface{}) {
	l.log(warnLevel, data...)
}

// Error logs at the error level
func (l *Logger) Error(data ...interface{}) {
	l.log(errorLevel, data...)
}

// log logs a message at the specified level
func (l *Logger) log(level *level, data ...interface{}) {

	if l.level.value > level.value {
		return
	}

	now := time.Now().Format(time.RFC3339)
	message := ""
	for _, item := range data {
		message += fmt.Sprint(item, " ")
	}

	l.Logpher.Configuration.writer.write(level.colourizer("[%s] [%s] [%s] %s", now, l.name, level.display, message))
}

// NewLogger creates a new logger using the autumn Logpher instance configuration
func NewLogger(name string) *Logger {
	return &Logger{name: name}
}

// GetLeafName gets the autumn leaf name
func (l *Logger) GetLeafName() string {
	return strings.ToLower(l.name) + "Logger"
}

// PostConstruct initializes the logger when it's used as an autumn leaf
func (l *Logger) PostConstruct() {
	l.level = newLevel(l.Logpher.Configuration.getLevel(l.name))
	l.name = strings.ToUpper(l.name)
}
