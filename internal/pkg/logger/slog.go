package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

type slogLogger struct {
	log *slog.Logger
}

func SetSlogLogger() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog := slogLogger{log: log}

	SetLogger(slog)
}

func (l slogLogger) Error(args ...interface{}) {
	l.log.Error(fmt.Sprint(args...))
}
func (l slogLogger) Fatal(args ...interface{}) {
	l.log.Error(fmt.Sprint(args...))
}
func (l slogLogger) Errorf(format string, args ...interface{}) {
	l.log.Error(fmt.Sprintf(format, args...))
}
func (l slogLogger) Info(args ...interface{}) {
	l.log.Info(fmt.Sprint(args...))
}
func (l slogLogger) Infof(format string, args ...interface{}) {
	l.log.Info(fmt.Sprintf(format, args...))
}
func (l slogLogger) Warn(args ...interface{}) {
	l.log.Warn(fmt.Sprint(args...))
}
func (l slogLogger) Warnf(format string, args ...interface{}) {
	l.log.Warn(fmt.Sprintf(format, args...))
}
func (l slogLogger) Debug(args ...interface{}) {
	l.log.Debug(fmt.Sprint(args...))
}
func (l slogLogger) Debugf(format string, args ...interface{}) {
	l.log.Debug(fmt.Sprintf(format, args...))
}
func (l slogLogger) LogAttrs(ctx context.Context, level slog.Level, msg string, attrs ...slog.Attr) {
	l.log.LogAttrs(ctx, slog.Level(level), msg, attrs...)
}
