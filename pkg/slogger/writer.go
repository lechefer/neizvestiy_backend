package slogger

import (
	"errors"
	"go.uber.org/zap"
	"io"
	"sync"

	"go.uber.org/zap/zapcore"
)

var (
	_plLevelEnabler = zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l == zapcore.DebugLevel || l == zapcore.ErrorLevel
	})
	_plStacktraceEnabler = zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l == zapcore.ErrorLevel
	})
)

type writeSyncer struct {
	w  io.Writer
	mx sync.RWMutex
}

func newWriteSyncer(w io.Writer) zapcore.WriteSyncer {
	return &writeSyncer{
		w: w,
	}
}

func (s *writeSyncer) Write(p []byte) (int, error) {
	s.mx.Lock()
	defer s.mx.Unlock()
	return s.w.Write(p)
}

func (s *writeSyncer) Sync() error {
	return errors.New("not implemented")
}
