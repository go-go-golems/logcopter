package logcopter

import (
	"io"
	"testing"

	"github.com/rs/zerolog"
)

func BenchmarkLoggerLookupDisabledLevel(b *testing.B) {
	m := NewManager()
	log := m.Package("app.benchmark.lookup")
	if err := m.Configure(zerolog.New(io.Discard), Config{Level: "info"}); err != nil {
		b.Fatalf("Configure returned error: %v", err)
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		log.Debug().Msg("dropped")
	}
}

func BenchmarkLoggerLookupEnabledLevel(b *testing.B) {
	m := NewManager()
	log := m.Package("app.benchmark.lookup")
	if err := m.Configure(zerolog.New(io.Discard), Config{Level: "debug"}); err != nil {
		b.Fatalf("Configure returned error: %v", err)
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		log.Debug().Msg("emitted")
	}
}
