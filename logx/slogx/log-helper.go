package slogx

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"time"
)

// LogHelper wraps a slog.Logger and provides helper methods like formatted logging.
type LogHelper struct {
	*slog.Logger
}

// log logs a message at the given level.
// It retrieves the caller information to include in the log record, so the caller of LogHelper will be included in logs (not the called LogHelper method).
// "Designed" to be called by other LogHelper methods like Fatal.
func (lh LogHelper) log(level slog.Level, msg string, args ...any) {
	if !lh.Enabled(context.Background(), level) {
		return
	}

	var pcs [1]uintptr
	runtime.Callers(3, pcs[:]) // skip Callers + log + caller LogHelper method
	r := slog.NewRecord(time.Now(), level, msg, pcs[0])
	r.Add(args...)

	_ = lh.Handler().Handle(context.Background(), r)
}

// logf logs a formatted message at the given level.
// It retrieves the caller information to include in the log record, so the caller of LogHelper will be included in logs (not the called LogHelper method).
// "Designed" to be called by other LogHelper methods like Warnf.
func (lh LogHelper) logf(level slog.Level, format string, args ...any) {
	if !lh.Enabled(context.Background(), level) {
		return
	}

	var pcs [1]uintptr
	runtime.Callers(3, pcs[:]) // skip Callers + logf + caller LogHelper method
	r := slog.NewRecord(time.Now(), level, fmt.Sprintf(format, args...), pcs[0])

	_ = lh.Handler().Handle(context.Background(), r)
}

// Debugf logs a formatted debug message.
func (lh LogHelper) Debugf(format string, a ...any) {
	lh.logf(slog.LevelDebug, format, a...)
}

// Infof logs a formatted info message.
func (lh LogHelper) Infof(format string, a ...any) {
	lh.logf(slog.LevelInfo, format, a...)
}

// Warnf logs a formatted warning message.
func (lh LogHelper) Warnf(format string, a ...any) {
	lh.logf(slog.LevelWarn, format, a...)
}

// Errorf logs a formatted error message.
func (lh LogHelper) Errorf(format string, a ...any) {
	lh.logf(slog.LevelError, format, a...)
}

// Fatalf logs a formatted error message and calls [os.Exit](1).
func (lh LogHelper) Fatalf(format string, a ...any) {
	lh.logf(slog.LevelError, format, a...)
	os.Exit(1)
}

// Fatal is equivalent to [log/slog.Logger.Error] followed by [os.Exit](1).
func (lh LogHelper) Fatal(msg string, args ...any) {
	lh.log(slog.LevelError, msg, args...)
	os.Exit(1)
}

// With returns a new LogHelper wrapping a new logger returned by the wrapped logger's With() method.
func (lh LogHelper) With(args ...any) LogHelper {
	return LogHelper{Logger: lh.Logger.With(args...)}
}

// Group returns a new LogHelper wrapping a new logger returned by the wrapped logger's WithGroup() method.
func (lh LogHelper) WithGroup(name string) LogHelper {
	return LogHelper{Logger: lh.Logger.WithGroup(name)}
}
