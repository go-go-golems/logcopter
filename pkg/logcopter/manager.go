package logcopter

import (
	"sort"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

const defaultAreaField = "area"

type Manager struct {
	mu    sync.Mutex
	state atomic.Value // stores *state
}

type state struct {
	base         zerolog.Logger
	defaultLevel zerolog.Level
	configured   map[string]zerolog.Level
	known        map[string]struct{}
	strictAreas  bool
	areaField    string
}

func NewManager() *Manager {
	m := &Manager{}
	m.state.Store(newDefaultState())
	return m
}

func newDefaultState() *state {
	return &state{
		base:         zerolog.Nop(),
		defaultLevel: zerolog.Disabled,
		configured:   map[string]zerolog.Level{},
		known:        map[string]struct{}{},
		areaField:    defaultAreaField,
	}
}

func (m *Manager) Configure(base zerolog.Logger, cfg Config) error {
	if m == nil {
		return errors.Errorf("nil logcopter manager")
	}

	defaultLevelName := strings.TrimSpace(cfg.Level)
	if defaultLevelName == "" {
		defaultLevelName = "info"
	}
	defaultLevel, err := ParseLevel(defaultLevelName)
	if err != nil {
		return errors.Wrap(err, "parse default logging level")
	}

	configured := make(map[string]zerolog.Level, len(cfg.Areas))
	for area, levelName := range cfg.Areas {
		normalizedArea, err := NormalizeArea(area)
		if err != nil {
			return errors.Wrapf(err, "invalid logging area %q", area)
		}
		level, err := ParseLevel(levelName)
		if err != nil {
			return errors.Wrapf(err, "invalid logging level for area %q", normalizedArea)
		}
		configured[normalizedArea] = level
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	old := m.currentState()
	next := &state{
		base:         base,
		defaultLevel: defaultLevel,
		configured:   configured,
		known:        cloneStringSet(old.known),
		strictAreas:  cfg.StrictAreas,
		areaField:    defaultAreaField,
	}

	if cfg.StrictAreas {
		warnings := validateConfiguredAreas(configured, next.known)
		if len(warnings) > 0 {
			return errors.Errorf("unknown logging area %q", warnings[0].Area)
		}
	}

	m.state.Store(next)
	return nil
}

func (m *Manager) Package(area string) Logger {
	return m.For(area)
}

func (m *Manager) For(area string) Logger {
	if m == nil {
		m = DefaultManager()
	}
	normalizedArea, err := NormalizeArea(area)
	if err != nil {
		return Logger{manager: m, area: strings.TrimSpace(area), invalid: err}
	}
	m.registerKnownArea(normalizedArea)
	return Logger{manager: m, area: normalizedArea}
}

func (m *Manager) Areas() []string {
	st := m.currentState()
	areas := make([]string, 0, len(st.known))
	for area := range st.known {
		areas = append(areas, area)
	}
	sort.Strings(areas)
	return areas
}

func (m *Manager) EffectiveLevel(area string) zerolog.Level {
	normalizedArea, err := NormalizeArea(area)
	if err != nil {
		return zerolog.Disabled
	}
	st := m.currentState()
	return resolveLevel(normalizedArea, st.configured, st.defaultLevel)
}

func (m *Manager) ValidateAreas(strict bool) []AreaWarning {
	st := m.currentState()
	warnings := validateConfiguredAreas(st.configured, st.known)
	if !strict {
		return warnings
	}
	return warnings
}

func (m *Manager) resolve(area string) zerolog.Logger {
	st := m.currentState()
	level := resolveLevel(area, st.configured, st.defaultLevel)
	return st.base.With().Str(st.areaField, area).Logger().Level(level)
}

func (m *Manager) currentState() *state {
	if m == nil {
		return newDefaultState()
	}
	st, ok := m.state.Load().(*state)
	if !ok || st == nil {
		return newDefaultState()
	}
	return st
}

func (m *Manager) registerKnownArea(area string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	old := m.currentState()
	if _, ok := old.known[area]; ok {
		return
	}

	next := old.clone()
	next.known[area] = struct{}{}
	m.state.Store(next)
}

func (s *state) clone() *state {
	return &state{
		base:         s.base,
		defaultLevel: s.defaultLevel,
		configured:   cloneLevelMap(s.configured),
		known:        cloneStringSet(s.known),
		strictAreas:  s.strictAreas,
		areaField:    s.areaField,
	}
}

func cloneLevelMap(in map[string]zerolog.Level) map[string]zerolog.Level {
	out := make(map[string]zerolog.Level, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
}

func cloneStringSet(in map[string]struct{}) map[string]struct{} {
	out := make(map[string]struct{}, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
}

func validateConfiguredAreas(configured map[string]zerolog.Level, known map[string]struct{}) []AreaWarning {
	warnings := []AreaWarning{}
	for area := range configured {
		if configuredAreaMatchesKnown(area, known) {
			continue
		}
		warnings = append(warnings, AreaWarning{
			Area:    area,
			Message: "configured logging area does not match any known generated/registered area",
		})
	}
	sort.Slice(warnings, func(i, j int) bool { return warnings[i].Area < warnings[j].Area })
	return warnings
}

func configuredAreaMatchesKnown(configuredArea string, known map[string]struct{}) bool {
	for area := range known {
		if area == configuredArea || strings.HasPrefix(area, configuredArea+".") || strings.HasPrefix(configuredArea, area+".") {
			return true
		}
	}
	return false
}
