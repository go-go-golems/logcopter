package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "logcopter-gen: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	cfg := generatorConfig{}
	fs := flag.NewFlagSet("logcopter-gen", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	fs.StringVar(&cfg.LogcopterImport, "logcopter-import", "github.com/go-go-golems/logcopter/pkg/logcopter", "import path for the logcopter runtime package")
	fs.StringVar(&cfg.AreaPrefix, "area-prefix", "", "area prefix prepended to generated package areas")
	fs.StringVar(&cfg.StripPrefix, "strip-prefix", "", "import path prefix removed before converting package path to area")
	fs.StringVar(&cfg.Out, "out", "logcopter.go", "output filename")
	fs.StringVar(&cfg.VarName, "var", "log", "generated package logger variable name")
	fs.BoolVar(&cfg.IncludeMain, "include-main", false, "generate loggers for package main")
	fs.StringVar(&cfg.AreasOut, "areas-out", "", "optional generated registry file containing all discovered areas")
	fs.BoolVar(&cfg.Check, "check", false, "fail if generated files are not current")
	fs.BoolVar(&cfg.DryRun, "dry-run", false, "print planned writes without changing files")

	fs.Usage = func() {
		_, _ = fmt.Fprintf(fs.Output(), "Usage: logcopter-gen [flags] <packages...>\n\n")
		fs.PrintDefaults()
	}

	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return nil
		}
		return err
	}
	cfg.Patterns = fs.Args()
	return generate(cfg)
}
