package main

import (
	"fmt"
	"os"

	"github.com/go-go-golems/glazed/pkg/help"
	help_cmd "github.com/go-go-golems/glazed/pkg/help/cmd"
	logcopterdoc "github.com/go-go-golems/logcopter/pkg/doc"
	"github.com/spf13/cobra"
)

func main() {
	if err := newRootCommand().Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "logcopter-gen: %v\n", err)
		os.Exit(1)
	}
}

func newRootCommand() *cobra.Command {
	cfg := generatorConfig{}
	cmd := &cobra.Command{
		Use:   "logcopter-gen [flags] <packages...>",
		Short: "Generate package-scoped logcopter logger variables",
		Long: `logcopter-gen scans Go packages and generates package-level logger variables.

The generated files provide consistent log area names and make it possible for applications to configure log output per package area. Use --check in CI to verify generated files are current without modifying the working tree.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg.Patterns = args
			return generate(cfg)
		},
	}
	cmd.Flags().StringVar(&cfg.LogcopterImport, "logcopter-import", "github.com/go-go-golems/logcopter/pkg/logcopter", "import path for the logcopter runtime package")
	cmd.Flags().StringVar(&cfg.AreaPrefix, "area-prefix", "", "area prefix prepended to generated package areas")
	cmd.Flags().StringVar(&cfg.StripPrefix, "strip-prefix", "", "import path prefix removed before converting package path to area")
	cmd.Flags().StringVar(&cfg.Out, "out", "logcopter.go", "output filename")
	cmd.Flags().StringVar(&cfg.VarName, "var", "log", "generated package logger variable name")
	cmd.Flags().BoolVar(&cfg.IncludeMain, "include-main", false, "generate loggers for package main")
	cmd.Flags().StringVar(&cfg.AreasOut, "areas-out", "", "optional generated registry file containing all discovered areas")
	cmd.Flags().BoolVar(&cfg.Check, "check", false, "fail if generated files are not current")
	cmd.Flags().BoolVar(&cfg.DryRun, "dry-run", false, "print planned writes without changing files")

	helpSystem := help.NewHelpSystem()
	cobra.CheckErr(logcopterdoc.AddDocToHelpSystem(helpSystem))
	help_cmd.SetupCobraRootCommand(helpSystem, cmd)

	return cmd
}
