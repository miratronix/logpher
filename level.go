package logpher

import (
	"github.com/fatih/color"
	"strings"
)

const (
	traceString = "TRACE"
	debugString = "DEBUG"
	infoString  = "INFO"
	warnString  = "WARN"
	errString   = "ERROR"
	offString   = "OFF"
)

var (
	Trace = &level{0, traceString, color.WhiteString}
	Debug = &level{1, debugString, color.BlueString}
	Info  = &level{2, infoString, color.CyanString}
	Warn  = &level{3, warnString, color.YellowString}
	Error = &level{4, errString, color.RedString}
	Off   = &level{5, offString, nil}
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
	case traceString:
		return Trace
	case debugString:
		return Debug
	case infoString:
		return Info
	case warnString:
		return Warn
	case errString:
		return Error
	case offString:
		return Off
	default:
		return Info
	}
}
