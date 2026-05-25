package logcopter

import (
	"bytes"
	"strings"
	"testing"

	"github.com/rs/zerolog"
)

func TestWriterForOutputDefaultsToStderr(t *testing.T) {
	if WriterForOutput("") == nil {
		t.Fatalf("expected default writer")
	}
	if WriterForOutput("stderr") == nil {
		t.Fatalf("expected stderr writer")
	}
	if WriterForOutput("stdout") == nil {
		t.Fatalf("expected stdout writer")
	}
}

func TestWriterForFormatJSONKeepsWriter(t *testing.T) {
	buf := &bytes.Buffer{}
	writer := WriterForFormat(buf, OutputConfig{Format: "json"})
	if writer != buf {
		t.Fatalf("expected JSON format to return original writer")
	}
}

func TestWriterForFormatTextUsesConsoleWriter(t *testing.T) {
	buf := &bytes.Buffer{}
	writer := WriterForFormat(buf, OutputConfig{Format: "text", NoColor: true})
	logger := zerolog.New(writer)
	logger.Info().Msg("hello")
	if !strings.Contains(buf.String(), "hello") {
		t.Fatalf("expected console output to contain message, got %q", buf.String())
	}
}

func TestNewLoggerJSON(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := zerolog.New(WriterForFormat(buf, OutputConfig{Format: "json"}))
	logger.Info().Str("area", "app.test").Msg("hello")
	out := buf.String()
	if !strings.Contains(out, `"message":"hello"`) {
		t.Fatalf("expected JSON output to contain message, got %q", out)
	}
	if !strings.Contains(out, `"area":"app.test"`) {
		t.Fatalf("expected JSON output to contain area, got %q", out)
	}
}
