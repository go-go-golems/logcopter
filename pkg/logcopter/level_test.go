package logcopter

import (
	"testing"

	"github.com/rs/zerolog"
)

func TestParseLevel(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want zerolog.Level
	}{
		{name: "trace", in: "trace", want: zerolog.TraceLevel},
		{name: "debug uppercase", in: "DEBUG", want: zerolog.DebugLevel},
		{name: "info with spaces", in: " info ", want: zerolog.InfoLevel},
		{name: "warning alias", in: "warning", want: zerolog.WarnLevel},
		{name: "warn", in: "warn", want: zerolog.WarnLevel},
		{name: "error", in: "error", want: zerolog.ErrorLevel},
		{name: "fatal", in: "fatal", want: zerolog.FatalLevel},
		{name: "panic", in: "panic", want: zerolog.PanicLevel},
		{name: "off alias", in: "off", want: zerolog.Disabled},
		{name: "none alias", in: "none", want: zerolog.Disabled},
		{name: "disabled alias", in: "disabled", want: zerolog.Disabled},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseLevel(tt.in)
			if err != nil {
				t.Fatalf("ParseLevel(%q) returned error: %v", tt.in, err)
			}
			if got != tt.want {
				t.Fatalf("ParseLevel(%q) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestParseLevelInvalid(t *testing.T) {
	for _, in := range []string{"", " ", "verbose", "debuggy"} {
		t.Run(in, func(t *testing.T) {
			if _, err := ParseLevel(in); err == nil {
				t.Fatalf("ParseLevel(%q) returned nil error", in)
			}
		})
	}
}
