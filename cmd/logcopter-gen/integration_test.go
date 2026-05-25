package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestGeneratedSourceCompilesInTempModule(t *testing.T) {
	tmp := t.TempDir()
	pkgDir := filepath.Join(tmp, "view", "render")
	if err := os.MkdirAll(pkgDir, 0o755); err != nil {
		t.Fatalf("mkdir package: %v", err)
	}
	if err := os.WriteFile(filepath.Join(tmp, "go.mod"), []byte(`module example.com/app

go 1.25.0

require github.com/go-go-golems/logcopter v0.0.0

replace github.com/go-go-golems/logcopter => `+repoRoot(t)+`
`), 0o644); err != nil {
		t.Fatalf("write go.mod: %v", err)
	}
	if err := os.WriteFile(filepath.Join(pkgDir, "render.go"), []byte("package render\n\nfunc Render() {}\n"), 0o644); err != nil {
		t.Fatalf("write package: %v", err)
	}

	content, err := renderPackageLogger(packagePlan{
		PackageName:     "render",
		OutputPath:      filepath.Join(pkgDir, "logcopter.go"),
		LogcopterImport: "github.com/go-go-golems/logcopter/pkg/logcopter",
		VarName:         "log",
		Area:            "app.view.render",
	})
	if err != nil {
		t.Fatalf("renderPackageLogger returned error: %v", err)
	}
	if err := os.WriteFile(filepath.Join(pkgDir, "logcopter.go"), content, 0o644); err != nil {
		t.Fatalf("write generated source: %v", err)
	}

	if out, err := runGo(tmp, "mod", "tidy"); err != nil {
		t.Fatalf("go mod tidy failed: %v\n%s", err, out)
	}
	if out, err := runGo(tmp, "test", "./..."); err != nil {
		t.Fatalf("generated module did not compile: %v\n%s", err, out)
	}
}

func runGo(dir string, args ...string) ([]byte, error) {
	cmd := exec.Command("go", args...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "GOWORK=off")
	return cmd.CombinedOutput()
}

func repoRoot(t *testing.T) string {
	t.Helper()
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	root := filepath.Clean(filepath.Join(wd, "..", ".."))
	return filepath.ToSlash(root)
}
