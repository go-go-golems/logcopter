package logcopter

import (
	"testing"

	"github.com/rs/zerolog"
)

func TestNormalizeArea(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{name: "simple", in: "app", want: "app"},
		{name: "nested", in: "app.view.render", want: "app.view.render"},
		{name: "trim", in: " lib.protocol.rx ", want: "lib.protocol.rx"},
		{name: "hyphen underscore", in: "github.com.acme.my-lib.parser_v2", want: "github.com.acme.my-lib.parser_v2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NormalizeArea(tt.in)
			if err != nil {
				t.Fatalf("NormalizeArea(%q) returned error: %v", tt.in, err)
			}
			if got != tt.want {
				t.Fatalf("NormalizeArea(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestNormalizeAreaInvalid(t *testing.T) {
	for _, in := range []string{"", " ", ".app", "app.", "app..db", "app/db", "app db", "app:db"} {
		t.Run(in, func(t *testing.T) {
			if _, err := NormalizeArea(in); err == nil {
				t.Fatalf("NormalizeArea(%q) returned nil error", in)
			}
		})
	}
}

func TestResolveLevel(t *testing.T) {
	configured := map[string]zerolog.Level{
		"app.view":        zerolog.DebugLevel,
		"app.view.render": zerolog.TraceLevel,
		"app.db":          zerolog.WarnLevel,
	}

	tests := []struct {
		area string
		want zerolog.Level
	}{
		{area: "app.view.render.partial", want: zerolog.TraceLevel},
		{area: "app.view.assets", want: zerolog.DebugLevel},
		{area: "app.db.sql", want: zerolog.WarnLevel},
		{area: "app.http", want: zerolog.InfoLevel},
	}

	for _, tt := range tests {
		t.Run(tt.area, func(t *testing.T) {
			got := resolveLevel(tt.area, configured, zerolog.InfoLevel)
			if got != tt.want {
				t.Fatalf("resolveLevel(%q) = %v, want %v", tt.area, got, tt.want)
			}
		})
	}
}
