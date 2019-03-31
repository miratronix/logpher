package logpher

import (
	"fmt"
	"sync"
)

// consoleWriter defines a basic console based writer
type consoleWriter struct {
	lock   *sync.Mutex
	closed bool
}

// newConsoleWriter creates a new console based writer
func newConsoleWriter() *consoleWriter {
	return &consoleWriter{
		lock: &sync.Mutex{},
	}
}

// write writes a log line to the console
func (c *consoleWriter) write(line string) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if !c.closed {
		fmt.Println(line)
	}
}

// close closes the writer
func (c *consoleWriter) close() {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.closed = true
}
