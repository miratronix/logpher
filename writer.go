package logpher

const (
	console = "console"
	file    = "file"
	rolling = "rolling"
)

// writer defines a basic log writer interface
type writer interface {
	colourEnabled() bool
	write(line string)
	close()
}
