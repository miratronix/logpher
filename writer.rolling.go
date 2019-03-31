package logpher

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// rollingWriter defines a log writer that rotates files up to the maximum count
type rollingWriter struct {
	lock         *sync.Mutex
	closed       bool
	file         *os.File
	fileName     string
	maxSize      int64
	maxCount     int
	bytesWritten int64
}

// newRollingWriter creates a new rolling writer
func newRollingWriter(fileName string, maxSize int, maxCount int) *rollingWriter {
	writer := &rollingWriter{
		lock:         &sync.Mutex{},
		file:         nil,
		fileName:     toAbsolutePath(fileName),
		maxSize:      int64(maxSize * megabyte),
		maxCount:     maxCount,
		bytesWritten: 0,
	}

	// Check if there's already a live log file
	info, err := os.Stat(writer.fileName)
	if err != nil {

		// Make sure it's a file doesn't exist error
		if !os.IsNotExist(err) {
			panic(err)
		}

		// Create the live file
		writer.file, err = openFile(writer.fileName)
		panicOnError(err)

		// Delete old files
		panicOnError(writer.deleteOld())
		return writer
	}

	// The file already exists, open it up
	writer.file, err = openFile(writer.fileName)
	panicOnError(err)

	// Store the size of it and rotate if necessary
	writer.bytesWritten = info.Size()
	if writer.bytesWritten >= writer.maxSize {
		panicOnError(writer.rotate())
	}

	// Delete old files
	panicOnError(writer.deleteOld())
	return writer
}

// rotate renames the current live file and creates a new one
func (r *rollingWriter) rotate() error {

	// Close the open file
	err := r.file.Close()
	if err != nil {
		return err
	}

	// Rename it
	err = os.Rename(r.fileName, r.fileName+"-"+time.Now().Format(time.RFC3339))
	if err != nil {
		return err
	}

	// Create a new "live" file
	r.bytesWritten = 0
	r.file, err = openFile(r.fileName)
	return err
}

// deleteOld deletes old log files, based on the configured max count
func (r *rollingWriter) deleteOld() error {

	// Get the log directory
	directory := filepath.Dir(r.fileName)

	// Walk the directory and delete any extra files
	backupCount := 0
	return filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {

		// Not a matching file
		if !strings.HasPrefix(path, r.fileName) {
			return nil
		}

		// Matching file, check if it has a timestamp on the end
		split := strings.Split(path, "-")
		_, err = time.Parse(time.RFC3339, split[len(split)-1])
		if err != nil {
			return nil
		}

		// Delete the oldest backup files
		backupCount++
		if backupCount > r.maxCount {
			return os.Remove(path)
		}

		return nil
	})
}

// colourEnabled determines if colour output is enabled for this writer
func (r *rollingWriter) colourEnabled() bool {
	return false
}

// write writes a log line to the file
func (r *rollingWriter) write(line string) {
	r.lock.Lock()
	defer r.lock.Unlock()

	if r.closed {
		return
	}

	count, err := r.file.WriteString(line + "\n")
	if err != nil {
		fmt.Println("Failed to write log line:", err)
		return
	}

	// Rotate if we've written more than we're allowed in the file
	r.bytesWritten += int64(count)
	if r.bytesWritten >= r.maxSize {

		err := r.rotate()
		if err != nil {
			fmt.Println("Failed to rotate log file:", err)
		}

		err = r.deleteOld()
		if err != nil {
			fmt.Println("Failed to delete old log file:", err)
		}
	}
}

// close closes the writer
func (r *rollingWriter) close() {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.file.Close()
	r.closed = true
}
