package logpher

import "strings"

// Logpher defines the main logging structure
type Logpher struct {
	Configuration *Configuration `autumn:"logConfiguration"`
}

// New creates a new logpher instance with the supplied configuration
func New(configuration *Configuration) *Logpher {

	// Initialize with the default configuration
	l := &Logpher{
		Configuration: &Configuration{Levels: map[string]string{}},
	}

	// A proper configuration was supplied, use that
	if configuration != nil {
		l.Configuration = configuration
	}

	// Apply the configuration
	l.PostConstruct()
	return l
}

// NewLogger constructs the a new logger with the specified name
func (l *Logpher) NewLogger(name string) *Logger {
	return newLogger(name, l)
}

// Close closes the log writer
func (l *Logpher) Close() {
	l.Configuration.writer.close()
}

// GetLeafName gets the autumn leaf name
func (l *Logpher) GetLeafName() string {
	return "logpher"
}

// PostConstruct enables autumn post construct functionality
func (l *Logpher) PostConstruct() {
	l.Configuration.writer = l.createWriter(l.Configuration.Type, false)
}

// PreDestroy enables autumn pre destroy functionality
func (l *Logpher) PreDestroy() {
	l.Close()
}

// createWriter creates a writer with the supplied type
func (l *Logpher) createWriter(writerType string, recursive bool) writer {
	switch strings.ToLower(writerType) {
	case combination:

		// Prevent infinite recursion when a combination writer is set as a sub type of a combination writer
		if recursive {
			panic("a combination writer can only be used at the top level")
		}

		// Split the sub writer string
		subTypes := strings.Split(l.Configuration.Combine, combinationDelimiter)
		if len(subTypes) < 1 {
			panic("please supply some writers to combine")
		}

		// Create the sub writers recursively
		subWriters := make([]writer, len(subTypes))
		for i, subWriterType := range subTypes {
			subWriters[i] = l.createWriter(subWriterType, true)
		}

		return newCombinationWriter(subWriters)

	case file:
		return newFileWriter(l.Configuration.File)

	case rolling:
		return newRollingWriter(l.Configuration.File, l.Configuration.Size, l.Configuration.Count)

	case console:
		fallthrough
	default:
		return newConsoleWriter()
	}
}
