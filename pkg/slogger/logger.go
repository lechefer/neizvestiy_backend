package slogger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	WithSkip(s int) Logger

	// Debug записывает сообщение в лог в момент вызова
	Debug(message string)
	// Debugf форматирует сообщение и записывает его в момент вызова
	Debugf(format string, args ...interface{})

	// Error записывает сообщение в лог в момент вызова
	Error(message string)
	// Errorf форматирует сообщение и записывает его в момент вызова
	Errorf(format string, args ...interface{})
}

// CustomLogger реализация Logger основанная на zap.Logger
type CustomLogger struct {
	l        *zap.Logger
	addStack bool
}

// NewLogger возвращает новый экземпляр логгера
func NewLogger(appName string, options ...Option) CustomLogger {
	jsonConfig := zapcore.EncoderConfig{
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeTime:    zapcore.RFC3339TimeEncoder,
		EncodeCaller:  zapcore.ShortCallerEncoder,
		MessageKey:    "message",
		LevelKey:      "level",
		TimeKey:       "date",
		CallerKey:     "caller",
		StacktraceKey: "stackTrace",
	}

	optionsHolder := optionHolder{
		out: os.Stdout,
		err: os.Stderr,
	}
	for _, option := range options {
		optionsHolder = option.apply(optionsHolder)
	}

	skip := 1
	if optionsHolder.skip > 0 {
		skip += optionsHolder.skip
	}

	zapOptions := []zap.Option{
		zap.ErrorOutput(newWriteSyncer(optionsHolder.err)),
		zap.AddCaller(),
		zap.AddCallerSkip(skip),
	}

	encoder := zapcore.NewJSONEncoder(jsonConfig)
	if optionsHolder.development {
		zapOptions = append(zapOptions, zap.Development())
		encoder = zapcore.NewConsoleEncoder(jsonConfig)
	}
	containerId, _ := os.Hostname()

	log := zap.New(
		zapcore.NewCore(
			encoder,
			newWriteSyncer(optionsHolder.out),
			_plLevelEnabler),
		zapOptions...,
	).With(
		zap.String("appName", appName),
		zap.String("containerId", containerId),
	)

	logger := CustomLogger{
		l:        log,
		addStack: optionsHolder.stacktrace,
	}

	return logger
}

func (s CustomLogger) WithSkip(n int) Logger {
	s.l = s.l.WithOptions(zap.AddCallerSkip(n))
	return s
}

func (s CustomLogger) Debug(message string) {
	s.l.Debug(message)
}

func (s CustomLogger) Debugf(format string, args ...interface{}) {
	s.l.Debug(fmt.Sprintf(format, args...))
}

func (s CustomLogger) Error(message string) {
	s.l.Error(message)
}

func (s CustomLogger) Errorf(format string, args ...interface{}) {
	s.l.Error(fmt.Sprintf(format, args...))
}
