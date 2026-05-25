package logcopter

import "github.com/rs/zerolog"

type Logger struct {
	manager *Manager
	area    string
	invalid error
}

func (l Logger) Trace() *zerolog.Event {
	logger := l.raw()
	return logger.Trace()
}

func (l Logger) Debug() *zerolog.Event {
	logger := l.raw()
	return logger.Debug()
}

func (l Logger) Info() *zerolog.Event {
	logger := l.raw()
	return logger.Info()
}

func (l Logger) Warn() *zerolog.Event {
	logger := l.raw()
	return logger.Warn()
}

func (l Logger) Error() *zerolog.Event {
	logger := l.raw()
	return logger.Error()
}

func (l Logger) Fatal() *zerolog.Event {
	logger := l.raw()
	return logger.Fatal()
}

func (l Logger) Panic() *zerolog.Event {
	logger := l.raw()
	return logger.Panic()
}

func (l Logger) WithLevel(level zerolog.Level) *zerolog.Event {
	logger := l.raw()
	return logger.WithLevel(level)
}

// Raw returns the current underlying zerolog logger for this area.
//
// Captured raw loggers are not reload-aware. Prefer calling Debug/Info/etc. on
// Logger directly unless an extremely hot path needs to avoid manager lookup.
func (l Logger) Raw() zerolog.Logger {
	return l.raw()
}

func (l Logger) Area() string {
	return l.area
}

func (l Logger) IsZero() bool {
	return l.manager == nil && l.area == "" && l.invalid == nil
}

func (l Logger) raw() zerolog.Logger {
	if l.invalid != nil || l.area == "" {
		return zerolog.Nop()
	}
	m := l.manager
	if m == nil {
		m = DefaultManager()
	}
	return m.resolve(l.area)
}
