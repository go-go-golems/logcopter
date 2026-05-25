package logcopter

import (
	"bytes"
	"strings"
	"testing"

	"github.com/rs/zerolog"
)

func TestLoggerCreatedBeforeConfigureWorksAfterConfigure(t *testing.T) {
	m := NewManager()
	log := m.Package("app.view.render")

	log.Info().Msg("before configure")

	var buf bytes.Buffer
	if err := m.Configure(zerolog.New(&buf), Config{Level: "info"}); err != nil {
		t.Fatalf("Configure returned error: %v", err)
	}

	log.Info().Msg("after configure")

	got := buf.String()
	if !strings.Contains(got, "after configure") {
		t.Fatalf("expected configured logger output, got %q", got)
	}
	if !strings.Contains(got, `"area":"app.view.render"`) {
		t.Fatalf("expected area field, got %q", got)
	}
	if strings.Contains(got, "before configure") {
		t.Fatalf("pre-config log unexpectedly emitted: %q", got)
	}
}

func TestAreaLevelsAndReload(t *testing.T) {
	m := NewManager()
	view := m.Package("app.view.render")
	db := m.Package("app.db")

	var first bytes.Buffer
	if err := m.Configure(zerolog.New(&first), Config{
		Level: "info",
		Areas: map[string]string{
			"app.view": "debug",
			"app.db":   "warn",
		},
	}); err != nil {
		t.Fatalf("Configure returned error: %v", err)
	}

	view.Debug().Msg("view debug")
	db.Debug().Msg("db debug")
	db.Warn().Msg("db warn")

	if !strings.Contains(first.String(), "view debug") {
		t.Fatalf("expected view debug in first output: %q", first.String())
	}
	if strings.Contains(first.String(), "db debug") {
		t.Fatalf("did not expect db debug in first output: %q", first.String())
	}
	if !strings.Contains(first.String(), "db warn") {
		t.Fatalf("expected db warn in first output: %q", first.String())
	}

	var second bytes.Buffer
	if err := m.Configure(zerolog.New(&second), Config{
		Level: "info",
		Areas: map[string]string{
			"app.view.render": "trace",
			"app.db":          "error",
		},
	}); err != nil {
		t.Fatalf("reload Configure returned error: %v", err)
	}

	view.Trace().Msg("view trace")
	db.Warn().Msg("db warn after reload")
	db.Error().Msg("db error after reload")

	if !strings.Contains(second.String(), "view trace") {
		t.Fatalf("expected view trace after reload: %q", second.String())
	}
	if strings.Contains(second.String(), "db warn after reload") {
		t.Fatalf("did not expect db warn after reload: %q", second.String())
	}
	if !strings.Contains(second.String(), "db error after reload") {
		t.Fatalf("expected db error after reload: %q", second.String())
	}
}

func TestInvalidReloadKeepsPreviousState(t *testing.T) {
	m := NewManager()
	log := m.Package("app.view")

	var first bytes.Buffer
	if err := m.Configure(zerolog.New(&first), Config{Level: "debug"}); err != nil {
		t.Fatalf("Configure returned error: %v", err)
	}

	if err := m.Configure(zerolog.New(&bytes.Buffer{}), Config{Level: "verbose"}); err == nil {
		t.Fatalf("expected invalid reload error")
	}

	log.Debug().Msg("still debug")
	if !strings.Contains(first.String(), "still debug") {
		t.Fatalf("expected previous config to remain active, got %q", first.String())
	}
}

func TestAreasAndEffectiveLevel(t *testing.T) {
	m := NewManager()
	_ = m.Package("app.db.sql")
	_ = m.Package("app.view")

	areas := m.Areas()
	want := []string{"app.db.sql", "app.view"}
	if len(areas) != len(want) {
		t.Fatalf("Areas() = %v, want %v", areas, want)
	}
	for i := range want {
		if areas[i] != want[i] {
			t.Fatalf("Areas() = %v, want %v", areas, want)
		}
	}

	if err := m.Configure(zerolog.New(&bytes.Buffer{}), Config{Level: "info", Areas: map[string]string{"app.db": "warn"}}); err != nil {
		t.Fatalf("Configure returned error: %v", err)
	}
	if got := m.EffectiveLevel("app.db.sql"); got != zerolog.WarnLevel {
		t.Fatalf("EffectiveLevel(app.db.sql) = %v, want warn", got)
	}
	if got := m.EffectiveLevel("app.view"); got != zerolog.InfoLevel {
		t.Fatalf("EffectiveLevel(app.view) = %v, want info", got)
	}
}

func TestStrictAreaValidation(t *testing.T) {
	m := NewManager()
	_ = m.Package("app.view.render")

	if err := m.Configure(zerolog.New(&bytes.Buffer{}), Config{
		Level:       "info",
		StrictAreas: true,
		Areas:       map[string]string{"app.view": "debug"},
	}); err != nil {
		t.Fatalf("parent area should match known descendant: %v", err)
	}

	if err := m.Configure(zerolog.New(&bytes.Buffer{}), Config{
		Level:       "info",
		StrictAreas: true,
		Areas:       map[string]string{"app.unknown": "debug"},
	}); err == nil {
		t.Fatalf("expected strict unknown area error")
	}
}
