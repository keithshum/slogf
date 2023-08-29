// 
// slogf sets up a logger based on Go's slog package at 1.21.
//
// Idea is to switch to a structural logging that gives key and value at the same time.
//
// Specs:
//
// - Support 2 log levels, debug and info. Setting debug = true for debug level while false means
//   info level.
//   Debug level displays DEBUG, INFO, WARN, ERROR, FATAL logs.
//   Info level displays INFO, WARN, ERROR, FATAL logs.
//
// - Support 2 log formats, Text and JSON.
//   E.g. Text
//   time=2023-07-11T17:12:46.649Z level=INFO source=main.go:29 msg="Entered main."
//   E.g. JSON
//   {"time":"2023-07-11T17:05:15.924382Z","level":"INFO","source":{"function":"main.main","file":"main.go","line":29},"msg":"Entered main."}
//
// - Support 2 log styles
//   1. Extra key value pairs
//   Debug(), Info(), Warn(), Error(), Fatal() supports extra arguments and attributes.
//   Way to call:
//   Debug(message, key1, value1, key2, value2) while message is string type value.
//   For passing err directly, just append .Error() after err first!
//   E.g. Debug("Hello world!", "Hello", "Peter Parker")
//   =>
//   {"time":"2023-07-11T17:05:15.924556Z","level":"DEBUG","source":{"function":"main.main","file":"main.go","line":32},"msg":"Hello world!", "Hello":"Peter Parker"}
//   2. Print format one liner
//   Debugf(), Infof(), Warnf(), Errorf(), Fatalf() supports print format style.
//   Way to call:
//   Debugf(format, substitue)
//   E.g. Debugf("Hello, %v!", "Peter Parker")
//   =>
//   {"time":"2023-07-11T17:05:15.924556Z","level":"DEBUG","source":{"function":"main.main","file":"main.go","line":32},"msg":"Hello, Peter Parker!"}
//
//

package slogf

import (
	"strings"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
	"log/slog"
)

var (
	Logger *slog.Logger
)

const (
	LevelFatal	= slog.Level(12)
)

//
// Actual logging in different levels.
//
// Debug() wraps around slog.Debug()
func Debug(format string, args ...any) {
	if !Logger.Enabled(context.Background(), slog.LevelDebug) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
	r := slog.NewRecord(time.Now(), slog.LevelDebug, format, pcs[0])
	r.Add(args...)
	_ = Logger.Handler().Handle(context.Background(), r)
}
//
// Debugf() provides flexibility to log with the 'printf' style
func Debugf(format string, args ...any) {
	if !Logger.Enabled(context.Background(), slog.LevelDebug) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
	r := slog.NewRecord(time.Now(), slog.LevelDebug, fmt.Sprintf(format, args...), pcs[0])
	_ = Logger.Handler().Handle(context.Background(), r)
}
//
// Info() wraps around slog.Info()
func Info(format string, args ...any) {
	if !Logger.Enabled(context.Background(), slog.LevelInfo) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
	r := slog.NewRecord(time.Now(), slog.LevelInfo, format, pcs[0])
	r.Add(args...)
	_ = Logger.Handler().Handle(context.Background(), r)
}
//
// Infof() provides flexibility to log with the 'printf' style
func Infof(format string, args ...any) {
	if !Logger.Enabled(context.Background(), slog.LevelInfo) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
	r := slog.NewRecord(time.Now(), slog.LevelInfo, fmt.Sprintf(format, args...), pcs[0])
	_ = Logger.Handler().Handle(context.Background(), r)
}
//
// Warn() wraps around slog.Warn()
func Warn(format string, args ...any) {
	if !Logger.Enabled(context.Background(), slog.LevelWarn) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
	r := slog.NewRecord(time.Now(), slog.LevelWarn, format, pcs[0])
	r.Add(args...)
	_ = Logger.Handler().Handle(context.Background(), r)
}
//
// Warnf() provides flexibility to log with the 'printf' style
func Warnf(format string, args ...any) {
	if !Logger.Enabled(context.Background(), slog.LevelWarn) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
	r := slog.NewRecord(time.Now(), slog.LevelWarn, fmt.Sprintf(format, args...), pcs[0])
	_ = Logger.Handler().Handle(context.Background(), r)
}
//
// Error() wraps around slog.Error()
func Error(format string, args ...any) {
	if !Logger.Enabled(context.Background(), slog.LevelError) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
	r := slog.NewRecord(time.Now(), slog.LevelError, format, pcs[0])
	r.Add(args...)
	_ = Logger.Handler().Handle(context.Background(), r)
}
//
// Errorf() provides flexibility to log with the 'printf' style
func Errorf(format string, args ...any) {
	if !Logger.Enabled(context.Background(), slog.LevelError) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
	r := slog.NewRecord(time.Now(), slog.LevelError, fmt.Sprintf(format, args...), pcs[0])
	_ = Logger.Handler().Handle(context.Background(), r)
}
//
// Fatal() exits the main program.
func Fatal(format string, args ...any) {
//func Fatal(format string, args ...any) {
	if !Logger.Enabled(context.Background(), LevelFatal) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
	r := slog.NewRecord(time.Now(), LevelFatal, format, pcs[0])
	r.Add(args...)
	_ = Logger.Handler().Handle(context.Background(), r)
	os.Exit(1)
}
//
// Fatalf() provides flexibility to log with the 'printf' style
func Fatalf(format string, args ...any) {
	if !Logger.Enabled(context.Background(), LevelFatal) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
	r := slog.NewRecord(time.Now(), LevelFatal, fmt.Sprintf(format, args...), pcs[0])
	_ = Logger.Handler().Handle(context.Background(), r)
	os.Exit(1)
}

//
// debug = false: INFO level displays INFO, WARN, ERROR, FATAL logs.
// debug = true: DEBUG level displays DEBUG, INFO, WARN, ERROR, FATAL logs.
//
// InitLogging() wraps around a new global logger with level and format.
func InitLogging(debug bool, format string) {

	replace := func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.SourceKey {
			source := a.Value.Any().(*slog.Source)
			source.File = filepath.Base(source.File)
		}
		
		// Adding a whole new level as Fatal
		if a.Key == slog.LevelKey {
			a.Key = "level"
			level := a.Value.Any().(slog.Level)
			if level == LevelFatal {
				a.Value = slog.StringValue("FATAL")
			}
		}
		return a
	}

	if debug == true {
		if strings.ToLower(format) == "text" {
			Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug, ReplaceAttr: replace}))
		} else {
			Logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug, ReplaceAttr: replace}))
		}
	} else {
		if strings.ToLower(format) == "text" {
			Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: slog.LevelInfo, ReplaceAttr: replace}))
		} else {
			Logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: slog.LevelInfo, ReplaceAttr: replace}))
		}		
	}
}
