# Changelog

## Unreleased

### Added

- Initial zerolog-backed logcopter runtime with reload-aware package loggers.
- Hierarchical area-level configuration with longest-prefix matching.
- `logcopter-gen` for generated package-local logger files.
- Lightweight output helpers for stderr/stdout and JSON/text zerolog output.
- Markdown-only help entries that can be loaded by the Glaze binary without adding a Glazed dependency to logcopter.
- Examples for direct runtime use and reusable library area prefixes.

### Integration

- Glazed integration lives in `github.com/go-go-golems/glazed/pkg/cmds/logging`.
- Glazed supports `--log-config`, `--log-area`, `--strict-log-areas`, and `logging.areas`.

### Migration note for Glazed logging users

Existing Glazed applications should keep their current imports and setup calls:

```go
import "github.com/go-go-golems/glazed/pkg/cmds/logging"
```

Applications can opt into package/area-specific diagnostics by generating
logcopter package loggers and then using existing Glazed logging setup. No
separate logcopter adapter package is required.
