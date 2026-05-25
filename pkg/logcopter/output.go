package logcopter

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

const (
	OutputStderr = "stderr"
	OutputStdout = "stdout"
	FormatJSON   = "json"
	FormatText   = "text"
)

// OutputConfig describes how to construct a base zerolog logger. It is kept
// intentionally small: application-framework concerns such as rotating files
// remain in callers such as Glazed.
type OutputConfig struct {
	Output     string
	Format     string
	WithCaller bool
	Timestamp  bool
	NoColor    bool
	TimeFormat string
}

// WriterForOutput returns the process stream selected by name. Empty output
// defaults to stderr.
func WriterForOutput(output string) io.Writer {
	switch strings.ToLower(strings.TrimSpace(output)) {
	case "", OutputStderr:
		return os.Stderr
	case OutputStdout:
		return os.Stdout
	default:
		return os.Stderr
	}
}

// ConsoleWriter wraps w in zerolog's text console writer with logcopter's
// default timestamp formatting.
func ConsoleWriter(w io.Writer, cfg OutputConfig) zerolog.ConsoleWriter {
	timeFormat := cfg.TimeFormat
	if timeFormat == "" {
		timeFormat = time.RFC3339Nano
	}
	return zerolog.ConsoleWriter{
		Out:        w,
		NoColor:    cfg.NoColor,
		TimeFormat: timeFormat,
	}
}

// WriterForFormat returns either the original writer for JSON logs or a console
// writer for text logs. Empty format defaults to text to match Glazed's CLI
// behavior.
func WriterForFormat(w io.Writer, cfg OutputConfig) io.Writer {
	switch strings.ToLower(strings.TrimSpace(cfg.Format)) {
	case FormatJSON:
		return w
	case "", FormatText:
		return ConsoleWriter(w, cfg)
	default:
		return ConsoleWriter(w, cfg)
	}
}

// NewLogger constructs a zerolog.Logger suitable for Manager.Configure. It does
// not set zerolog's global level; per-area filtering is owned by Manager.
func NewLogger(cfg OutputConfig) zerolog.Logger {
	writer := WriterForFormat(WriterForOutput(cfg.Output), cfg)
	logger := zerolog.New(writer)
	if cfg.Timestamp {
		logger = logger.With().Timestamp().Logger()
	}
	if cfg.WithCaller {
		logger = logger.With().Caller().Logger()
	}
	return logger
}
