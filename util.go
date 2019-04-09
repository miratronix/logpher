package logpher

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	megabyte = 1024 * 1024
	format   = "[%s] [%s] [%s] %s"
)

// panicOnError panics when a non-nil error is supplied
func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

// toAbsolutePath converts a file path to an absolute path, panicking if there are failures
func toAbsolutePath(path string) string {
	absolutePath, err := filepath.Abs(path)
	panicOnError(err)
	return absolutePath
}

// openFile opens the supplied file path
func openFile(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
}

// formatStandard formats a standard log line, without colouring it
func formatStandard(logger *Logger, level *level, line string) string {
	return fmt.Sprintf(format, time.Now().Format(time.RFC3339), logger.name, level.display, line)
}

// formatColour formats a log line with colour information
func formatColour(logger *Logger, level *level, line string) string {
	return level.colourizer(format, time.Now().Format(time.RFC3339), logger.name, level.display, line)
}
