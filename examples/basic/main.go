package main

import (
	"os"

	"github.com/go-go-golems/logcopter/pkg/logcopter"
)

var log = logcopter.Package("app.example.basic")

func main() {
	base := logcopter.NewLogger(logcopter.OutputConfig{
		Output:    logcopter.OutputStderr,
		Format:    logcopter.FormatText,
		Timestamp: true,
		NoColor:   true,
	})
	if err := logcopter.Configure(base, logcopter.Config{
		Level: "info",
		Areas: map[string]string{
			"app.example.basic": "debug",
		},
	}); err != nil {
		panic(err)
	}

	log.Debug().Str("pid", os.Args[0]).Msg("debug enabled for this area")
	log.Info().Msg("hello from logcopter")
}
