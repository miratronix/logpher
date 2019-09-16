package logpher

const defaultLevelKey = "default"

// Configuration defines the configuration structure for logging
type Configuration struct {
	Type    string // The main writer type
	Combine string // A comma separated string indicating which loggers to combine when using a combination writer
	File    string // The file path for file-based writers
	Size    int    // The maximum size in bytes for the rolling writer
	Count   int    // The maximum file count for the rolling writer
	Levels  map[string]string
	writer  writer
}

// NewConfiguration creates a new configuration object
func NewConfiguration() *Configuration {
	return &Configuration{
		Levels: map[string]string{},
	}
}

// getLevel gets the level for a logger
func (c *Configuration) getLevel(logger string) string {

	// No levels specified
	if c.Levels == nil {
		return infoString
	}

	// Check if the logger has an associated level
	level, ok := c.Levels[logger]
	if !ok {

		// Fall back to the configured default level
		defaultLevel, ok := c.Levels[defaultLevelKey]
		if ok {
			return defaultLevel
		}

		// No default level configured
		return infoString
	}

	return level
}
