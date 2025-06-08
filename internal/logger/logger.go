package logger

import (
	"log/slog"
	"os"
)

type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelSilent
)

type Logger interface {
	Debug(string, ...any)
	Info(string, ...any)
	Warn(string, ...any)
	Error(string, ...any)
	SetLevel(LogLevel)
}

type defaultLogger struct {
	lvl     LogLevel
	handler slog.Handler
}

func DefaultLogger(lvl ...LogLevel) Logger {

	logLevel := LevelInfo
	if len(lvl) > 0 {
		logLevel = lvl[0]
	}

	slogLevel := lvlToSloglvl(lvl[0])

	opts := &slog.HandlerOptions{
		Level: slogLevel,
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)

	return &defaultLogger{
		lvl:     logLevel,
		handler: handler,
	}

}

func lvlToSloglvl(lvl LogLevel) slog.Level {
	switch lvl {
	case LevelDebug:
		return slog.LevelDebug
	case LevelInfo:
		return slog.LevelInfo
	case LevelWarn:
		return slog.LevelWarn
	case LevelError:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func (l *defaultLogger) Debug(msg string, args ...any) {
	if l.lvl <= LevelDebug {
		slog.Debug(msg, args...)
	}
}

func (l *defaultLogger) Info(msg string, args ...any) {
	if l.lvl <= LevelInfo {
		slog.Info(msg, args...)
	}
}

func (l *defaultLogger) Warn(msg string, args ...any) {
	if l.lvl <= LevelWarn {
		slog.Warn(msg, args...)
	}
}

func (l *defaultLogger) Error(msg string, args ...any) {
	if l.lvl <= LevelError {
		slog.Error(msg, args...)
	}
}

func (l *defaultLogger) SetLevel(level LogLevel) {
	l.lvl = level
	slogLevel := lvlToSloglvl(level)

	opts := &slog.HandlerOptions{
		Level: slogLevel,
	}

	l.handler = slog.NewJSONHandler(os.Stdout, opts)
	slog.SetDefault(slog.New(l.handler))
}
