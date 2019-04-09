package logpher

import (
	"sync"
)

// combinationDelimiter defines the delimiter to use for combination writers
const combinationDelimiter = ","

// combinationWriter defines a simple writer that combines multiple writers into one
type combinationWriter struct {
	lock    *sync.Mutex
	closed  bool
	writers []writer
}

// newCombinationWriter creates a new combination writer
func newCombinationWriter(writers []writer) *combinationWriter {
	return &combinationWriter{
		lock:    &sync.Mutex{},
		writers: writers,
	}
}

// write writes a log line each underlying writer
func (c *combinationWriter) write(logger *Logger, level *level, line string) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.closed {
		return
	}

	for _, writer := range c.writers {
		writer.write(logger, level, line)
	}
}

// close closes the writer
func (c *combinationWriter) close() {
	c.lock.Lock()
	defer c.lock.Unlock()

	for _, writer := range c.writers {
		writer.close()
	}
	c.closed = true
}
