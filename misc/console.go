package misc

import (
	"fmt"
	"strings"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	TraceColor   = "\033[1;37m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[1;35m%s\033[0m"
)

type ConsoleLike interface {
	Info(format string, a ...interface{})
	Warn(format string, a ...interface{})
	Trace(format string, a ...interface{})
	Debug(format string, a ...interface{})
	Error(format string, a ...interface{})
}

type ColorfulConsole struct {
	ConsoleLike
}

func (console *ColorfulConsole) format(level string, format string) string {
	return strings.Replace(level, "%s", format, 1)
}

func (console *ColorfulConsole) Info(format string, a ...interface{}) {
	fmt.Printf(console.format(InfoColor, format), a...)
}

func (console *ColorfulConsole) Debug(format string, a ...interface{}) {
	fmt.Printf(console.format(DebugColor, format), a...)
}

func (console *ColorfulConsole) Trace(format string, a ...interface{}) {
	fmt.Printf(console.format(TraceColor, format), a...)
}

func (console *ColorfulConsole) Error(format string, a ...interface{}) {
	fmt.Printf(console.format(ErrorColor, format), a...)
}

func (console *ColorfulConsole) Warn(format string, a ...interface{}) {
	fmt.Printf(console.format(WarningColor, format), a...)
}

var Console = &ColorfulConsole{}
