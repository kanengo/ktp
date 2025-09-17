package logger

import (
	"log/slog"

	"github.com/go-kratos/kratos/v2/log"
)

type slogger struct {
	log *slog.Logger
}

// Log implements log.Logger.
func (sl *slogger) Log(level log.Level, keyvals ...any) error {
	switch level {
	case log.LevelDebug:
		sl.log.Debug("debug", keyvals...)
	case log.LevelInfo:
		sl.log.Info("info", keyvals...)
	case log.LevelWarn:
		sl.log.Warn("warn", keyvals...)
	case log.LevelError:
		sl.log.Error("error", keyvals...)
	default:
		sl.log.Debug("unknown level", keyvals...)
	}
	return nil
}

var _ log.Logger = (*slogger)(nil)
