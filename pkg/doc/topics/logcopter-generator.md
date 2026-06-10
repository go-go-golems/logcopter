---
Title: Logcopter Generator
Slug: logcopter-generator
Short: Generate package-scoped logcopter logger variables and area registries for Go packages.
Topics:
- logcopter
- logging
- code-generation
Commands:
- logcopter-gen
Flags:
- area-prefix
- strip-prefix
- out
- check
- dry-run
IsTopLevel: true
IsTemplate: false
ShowPerDefault: true
SectionType: GeneralTopic
---

`logcopter-gen` scans Go packages and generates package-level logger variables backed by the logcopter runtime. Rollouts usually generate a `logcopter.go` file in each package and optionally an area registry file that lets applications validate configured log areas.

A typical repository check looks like this:

```bash
logcopter-gen \
  --area-prefix go-go-golems.my-package \
  --strip-prefix github.com/go-go-golems/my-package \
  --check \
  ./...
```

Use `--dry-run` to inspect planned writes without changing files, and use `--check` in CI to fail when generated files are stale.
