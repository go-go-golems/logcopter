package main

import (
	"github.com/go-go-golems/logcopter/examples/library-prefix/decoder"
	"github.com/go-go-golems/logcopter/pkg/logcopter"
)

func main() {
	base := logcopter.NewLogger(logcopter.OutputConfig{
		Format:    logcopter.FormatText,
		Timestamp: true,
		NoColor:   true,
	})
	if err := logcopter.Configure(base, logcopter.Config{
		Level: "warn",
		Areas: map[string]string{
			"lib.ble": "trace",
		},
	}); err != nil {
		panic(err)
	}

	_ = decoder.Decode([]byte{0x01, 0x02, 0x03})
}
