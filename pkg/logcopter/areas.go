package logcopter

import (
	"strings"
	"unicode"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// NormalizeArea trims and validates an area name.
//
// Areas are dot-separated stable logging namespaces such as app.view.render or
// lib.protocol.parser. Each segment may contain letters, digits, underscores,
// and hyphens. Empty segments are rejected.
func NormalizeArea(area string) (string, error) {
	area = strings.TrimSpace(area)
	if area == "" {
		return "", errors.Errorf("empty logging area")
	}
	if strings.HasPrefix(area, ".") || strings.HasSuffix(area, ".") {
		return "", errors.Errorf("invalid logging area %q: leading or trailing dot", area)
	}

	parts := strings.Split(area, ".")
	for _, part := range parts {
		if part == "" {
			return "", errors.Errorf("invalid logging area %q: empty segment", area)
		}
		for _, r := range part {
			if isAreaRune(r) {
				continue
			}
			return "", errors.Errorf("invalid logging area %q: invalid character %q", area, r)
		}
	}

	return area, nil
}

func isAreaRune(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '-'
}

func resolveLevel(area string, configured map[string]zerolog.Level, fallback zerolog.Level) zerolog.Level {
	for current := area; current != ""; {
		if level, ok := configured[current]; ok {
			return level
		}

		i := strings.LastIndex(current, ".")
		if i < 0 {
			break
		}
		current = current[:i]
	}

	return fallback
}
