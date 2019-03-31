package logpher

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
	switch l.Configuration.Type {

	case file:
		l.Configuration.writer = newFileWriter(l.Configuration.File)

	case rolling:
		l.Configuration.writer = newRollingWriter(l.Configuration.File, l.Configuration.MaxSize, l.Configuration.MaxCount)

	case console:
		fallthrough
	default:
		l.Configuration.writer = newConsoleWriter()
	}
}

// PreDestroy enables autumn pre destroy functionality
func (l *Logpher) PreDestroy() {
	l.Close()
}
