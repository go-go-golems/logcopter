package main

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/tools/go/packages"
)

type generatorConfig struct {
	LogcopterImport string
	AreaPrefix      string
	StripPrefix     string
	Out             string
	VarName         string
	IncludeMain     bool
	AreasOut        string
	Check           bool
	DryRun          bool
	Patterns        []string
}

type packagePlan struct {
	PackageName     string
	PackagePath     string
	Directory       string
	OutputPath      string
	Area            string
	LogcopterImport string
	VarName         string
}

func planPackages(cfg generatorConfig) ([]packagePlan, error) {
	if len(cfg.Patterns) == 0 {
		return nil, fmt.Errorf("at least one package pattern is required")
	}
	if strings.TrimSpace(cfg.AreaPrefix) == "" {
		return nil, fmt.Errorf("-area-prefix is required")
	}
	if strings.TrimSpace(cfg.StripPrefix) == "" {
		return nil, fmt.Errorf("-strip-prefix is required")
	}
	if strings.TrimSpace(cfg.LogcopterImport) == "" {
		return nil, fmt.Errorf("-logcopter-import is required")
	}
	if strings.TrimSpace(cfg.VarName) == "" {
		return nil, fmt.Errorf("-var must not be empty")
	}
	if strings.TrimSpace(cfg.Out) == "" {
		return nil, fmt.Errorf("-out must not be empty")
	}

	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles,
	}, cfg.Patterns...)
	if err != nil {
		return nil, err
	}
	if packages.PrintErrors(pkgs) > 0 {
		return nil, fmt.Errorf("package loading failed")
	}

	plans := make([]packagePlan, 0, len(pkgs))
	for _, pkg := range pkgs {
		if pkg.Name == "main" && !cfg.IncludeMain {
			continue
		}
		dir, ok := packageDir(pkg)
		if !ok {
			continue
		}
		area, err := areaForPackage(pkg.PkgPath, cfg.StripPrefix, cfg.AreaPrefix)
		if err != nil {
			return nil, err
		}
		plans = append(plans, packagePlan{
			PackageName:     pkg.Name,
			PackagePath:     pkg.PkgPath,
			Directory:       dir,
			OutputPath:      filepath.Join(dir, cfg.Out),
			Area:            area,
			LogcopterImport: cfg.LogcopterImport,
			VarName:         cfg.VarName,
		})
	}

	sort.Slice(plans, func(i, j int) bool { return plans[i].PackagePath < plans[j].PackagePath })
	return plans, nil
}

func packageDir(pkg *packages.Package) (string, bool) {
	files := pkg.GoFiles
	if len(files) == 0 {
		files = pkg.CompiledGoFiles
	}
	if len(files) == 0 {
		return "", false
	}
	return filepath.Dir(files[0]), true
}

func areaForPackage(pkgPath, stripPrefix, areaPrefix string) (string, error) {
	pkgPath = strings.Trim(pkgPath, "/")
	stripPrefix = strings.Trim(stripPrefix, "/")
	areaPrefix = strings.Trim(areaPrefix, ".")

	if pkgPath == "" {
		return "", fmt.Errorf("empty package path")
	}
	if stripPrefix == "" {
		return "", fmt.Errorf("empty strip prefix")
	}
	if areaPrefix == "" {
		return "", fmt.Errorf("empty area prefix")
	}

	rel := ""
	switch {
	case pkgPath == stripPrefix:
		rel = ""
	case strings.HasPrefix(pkgPath, stripPrefix+"/"):
		rel = strings.TrimPrefix(pkgPath, stripPrefix+"/")
	default:
		return "", fmt.Errorf("package %q does not have strip prefix %q", pkgPath, stripPrefix)
	}

	if rel == "" {
		return areaPrefix, nil
	}
	return areaPrefix + "." + strings.ReplaceAll(rel, "/", "."), nil
}
