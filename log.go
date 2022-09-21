package colored_log

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

type ColoredLogger struct {
	ldef      *log.Logger
	lsucc     *log.Logger
	lerr      *log.Logger
	defColor  string
	succColor string
	errColor  string
	prefix    string
	flags     int
	out       io.Writer
	mu        sync.Mutex
}

func New(out io.Writer, prefix string, flag int) *ColoredLogger {
	return NewColored(out, prefix, Blue, Green, Red, flag)
}

func NewColored(out io.Writer, prefix string, default_color string, success_color string, error_color string, flag int) *ColoredLogger {
	return &ColoredLogger{
		ldef:      log.New(os.Stdout, fmt.Sprint(default_color, prefix), flag),
		lsucc:     log.New(os.Stdout, fmt.Sprint(success_color, prefix), flag),
		lerr:      log.New(os.Stdout, fmt.Sprint(error_color, prefix), flag),
		defColor:  default_color,
		succColor: success_color,
		errColor:  error_color,
		prefix:    prefix,
		flags:     flag,
		out:       out,
	}
}

var std_log = New(os.Stdout, "", log.LstdFlags)

func (l *ColoredLogger) Print(a ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.ldef.Print(a...)
}

func (l *ColoredLogger) Println(a ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.ldef.Println(a...)
}
func (l *ColoredLogger) Printf(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.ldef.Printf(format, v...)
}

func (l *ColoredLogger) Panic(a ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.lerr.Panic(a...)
}

func (l *ColoredLogger) Panicln(a ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.lerr.Panicln(a...)
}

func (l *ColoredLogger) Panicf(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.lerr.Panicf(format, v...)
}

func (l *ColoredLogger) Fatal(a ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.lerr.Panic(a...)
}

func (l *ColoredLogger) Fatalln(a ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.lerr.Fatalln(a...)
}

func (l *ColoredLogger) Fatalf(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.lerr.Fatalf(format, v...)
}

func (l *ColoredLogger) Success(a ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.lsucc.Print(a...)
}

func (l *ColoredLogger) Successln(a ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.lsucc.Println(a...)
}

func (l *ColoredLogger) Successf(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.lsucc.Printf(format, v...)
}

func (l *ColoredLogger) Flags() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.flags
}

// Sets log flag to input
func (l *ColoredLogger) SetFlags(i int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.ldef.SetFlags(i)
	l.lsucc.SetFlags(i)
	l.lerr.SetFlags(i)
}

func (l *ColoredLogger) Prefix() string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.prefix
}

// Sets prefix to input
func (l *ColoredLogger) SetPrefix(s string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.ldef.SetPrefix(fmt.Sprint(l.defColor, s))
	l.lsucc.SetPrefix(fmt.Sprint(l.succColor, s))
	l.lerr.SetPrefix(fmt.Sprint(l.errColor, s))
}

func (l *ColoredLogger) Writer() io.Writer {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.out
}

// Sets outputs of all types of logs to input wirter
func (l *ColoredLogger) SetOutput(w io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.ldef.SetOutput(w)
	l.lsucc.SetOutput(w)
	l.lerr.SetOutput(w)
}

func (l *ColoredLogger) Output(calldepth int, s string) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.ldef.Output(calldepth, s)
}

// Sets outputs of all types of logs to input wirter
func SetOutput(w io.Writer) {
	std_log.SetOutput(w)
}

// Flags returns the output flags for the standard logger.
// The flag bits are Ldate, Ltime, and so on.
func Flags() int {
	return std_log.Flags()
}

// Sets log flag to input
func SetFlags(flag int) {
	std_log.SetFlags(flag)
}

// Prefix returns the output prefix for the standard logger.
func Prefix() string {
	return std_log.Prefix()
}

// Sets prefix for all types of logger to input
func SetPrefix(prefix string) {
	std_log.SetPrefix(prefix)
}

// Writer returns the output destination for the standard logger.
func Writer() io.Writer {
	return std_log.Writer()
}

// Prints to io.Writer
// Uses log.Print, prints in default color
func Print(a ...interface{}) {
	std_log.Print(a...)
}

// Prints to io.Writer with new line
// Uses log.Println, prints in default color
func Println(a ...interface{}) {
	std_log.Println(a...)
}

// Prints to io.Writer with formated text
// Uses log.Printf, prints in default color
func Printf(format string, v ...interface{}) {
	std_log.Printf(format, v...)
}

// Similar to log.Panic
// Prints in error color
func Panic(a ...interface{}) {
	std_log.Panic(a...)
}

// Similar to log.Panicln
// Prints in error color
func Panicln(a ...interface{}) {
	std_log.Panicln(a...)
}

// Similar to log.Panicf
// Prints in error color
func Panicf(format string, v ...interface{}) {
	std_log.Panicf(format, v...)
}

// Similar to log.Fatal
// Prints in error color
func Fatal(a ...interface{}) {
	std_log.Panic(a...)
}

// Similar to log.Fataln
// Prints in error color
func Fatalln(a ...interface{}) {
	std_log.Fatalln(a...)
}

// Similar to log.Fatalf
// Prints in error color
func Fatalf(format string, v ...interface{}) {
	std_log.Fatalf(format, v...)
}

// Similar to Print
// Prints in success color
func Success(a ...interface{}) {
	std_log.Success(a...)
}

// Similar to Println
// Prints in success color
func Successln(a ...interface{}) {
	std_log.Successln(a...)
}

// Similar to Printf
// Prints in success color
func Successf(format string, v ...interface{}) {
	std_log.Successf(format, v...)
}

// Similar to logger Output()
func Output(calldepth int, s string) error {
	return std_log.Output(calldepth, s) // +1 for this frame.
}
