package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage: logcopter-gen [flags] <packages...>\n\n")
		flag.PrintDefaults()
	}

	if err := run(os.Args[1:]); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "logcopter-gen: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	fs := flag.NewFlagSet("logcopter-gen", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	logcopterImport := fs.String("logcopter-import", "github.com/go-go-golems/logcopter/pkg/logcopter", "import path for the logcopter runtime package")
	areaPrefix := fs.String("area-prefix", "", "area prefix prepended to generated package areas")
	stripPrefix := fs.String("strip-prefix", "", "import path prefix removed before converting package path to area")
	out := fs.String("out", "logcopter.go", "output filename")
	varName := fs.String("var", "log", "generated package logger variable name")
	includeMain := fs.Bool("include-main", false, "generate loggers for package main")
	areasOut := fs.String("areas-out", "", "optional generated registry file containing all discovered areas")
	check := fs.Bool("check", false, "fail if generated files are not current")
	dryRun := fs.Bool("dry-run", false, "print planned writes without changing files")

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
	_ = logcopterImport
	_ = areaPrefix
	_ = stripPrefix
	_ = out
	_ = varName
	_ = includeMain
	_ = areasOut
	_ = check
	_ = dryRun

	if fs.NArg() == 0 {
		fs.Usage()
		return fmt.Errorf("at least one package pattern is required")
	}

	return fmt.Errorf("generator implementation is not complete yet")
}
