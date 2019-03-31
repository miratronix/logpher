package logpher

import (
	"github.com/fatih/color"
	"strings"
)

const (
	trace = "TRACE"
	debug = "DEBUG"
	info  = "INFO"
	warn  = "WARN"
	err   = "ERROR"
)

var (
	traceLevel = &level{0, trace, color.WhiteString}
	debugLevel = &level{1, debug, color.BlueString}
	infoLevel  = &level{2, info, color.CyanString}
	warnLevel  = &level{3, warn, color.YellowString}
	errorLevel = &level{4, err, color.RedString}
)

// level defines a logging level
type level struct {
	value      int
	display    string
	colourizer func(format string, a ...interface{}) string
}

// newLevel constructs a new level from a string level name
func newLevel(level string) *level {
	switch strings.ToUpper(level) {
	case trace:
		return traceLevel
	case debug:
		return debugLevel
	case info:
		return infoLevel
	case warn:
		return warnLevel
	case err:
		return errorLevel
	default:
		return infoLevel
	}
}
