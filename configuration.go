package logpher

const defaultLevelKey = "default"

// Configuration defines the configuration structure for logging
type Configuration struct {
	Type     string
	File     string
	MaxSize  int
	MaxCount int
	Levels   map[string]string
	writer   writer
}

// getLevel gets the level for a logger
func (c *Configuration) getLevel(logger string) string {

	// No levels specified
	if c.Levels == nil {
		return info
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
		return info
	}

	return level
}
