package logpher

import (
	"os"
	"path/filepath"
)

const megabyte = 1024 * 1024

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
