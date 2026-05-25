package logcopter

import "github.com/rs/zerolog"

var defaultManager = NewManager()

func DefaultManager() *Manager {
	return defaultManager
}

func Configure(base zerolog.Logger, cfg Config) error {
	return defaultManager.Configure(base, cfg)
}

func Package(area string) Logger {
	return defaultManager.Package(area)
}

func For(area string) Logger {
	return defaultManager.For(area)
}

func Areas() []string {
	return defaultManager.Areas()
}

func EffectiveLevel(area string) zerolog.Level {
	return defaultManager.EffectiveLevel(area)
}

func ValidateAreas(strict bool) []AreaWarning {
	return defaultManager.ValidateAreas(strict)
}
