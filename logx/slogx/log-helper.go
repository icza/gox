package slogx

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

// LogHelper wraps a slog.Logger and provides helper methods like formatted logging.
type LogHelper struct {
	*slog.Logger
}

// Debugf logs a formatted debug message.
func (lh LogHelper) Debugf(format string, a ...any) {
	if lh.Enabled(context.Background(), slog.LevelDebug) {
		lh.Debug(fmt.Sprintf(format, a...))
	}
}

// Infof logs a formatted info message.
func (lh LogHelper) Infof(format string, a ...any) {
	if lh.Enabled(context.Background(), slog.LevelInfo) {
		lh.Info(fmt.Sprintf(format, a...))
	}
}

// Warnf logs a formatted warning message.
func (lh LogHelper) Warnf(format string, a ...any) {
	if lh.Enabled(context.Background(), slog.LevelWarn) {
		lh.Warn(fmt.Sprintf(format, a...))
	}
}

// Errorf logs a formatted error message.
func (lh LogHelper) Errorf(format string, a ...any) {
	if lh.Enabled(context.Background(), slog.LevelError) {
		lh.Error(fmt.Sprintf(format, a...))
	}
}

// Fatalf logs a formatted error message and exits the application.
func (lh LogHelper) Fatalf(format string, a ...any) {
	if lh.Enabled(context.Background(), slog.LevelError) {
		lh.Error(fmt.Sprintf(format, a...))
	}
	os.Exit(1)
}

// Fatal is equivalent to [log/slog.Logger.Error] followed by [os.Exit](1).
func (lh LogHelper) Fatal(msg string, args ...any) {
	lh.Error(msg, args...)
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
