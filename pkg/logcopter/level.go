package logcopter

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// ParseLevel parses a user-facing logging level string into a zerolog level.
//
// In addition to zerolog's normal level names, it accepts:
//   - warning as an alias for warn
//   - off, none, and disabled as aliases for zerolog.Disabled
func ParseLevel(s string) (zerolog.Level, error) {
	normalized := strings.ToLower(strings.TrimSpace(s))
	if normalized == "" {
		return zerolog.NoLevel, errors.Errorf("empty logging level")
	}

	switch normalized {
	case "off", "none", "disabled":
		return zerolog.Disabled, nil
	case "warning":
		return zerolog.WarnLevel, nil
	default:
		level, err := zerolog.ParseLevel(normalized)
		if err != nil {
			return zerolog.NoLevel, errors.Wrapf(err, "invalid logging level %q", s)
		}
		return level, nil
	}
}

func mustParseLevel(s string) zerolog.Level {
	level, err := ParseLevel(s)
	if err != nil {
		panic(err)
	}
	return level
}
