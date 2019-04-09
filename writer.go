package logpher

const (
	console     = "console"
	file        = "file"
	rolling     = "rolling"
	combination = "combination"
)

// writer defines a basic log writer interface
type writer interface {
	write(logger *Logger, level *level, line string)
	close()
}
