---
Title: ChatGPT proposal for package-scoped configurable logging
Ticket: LOGCOPTER-001
Status: active
Topics:
    - logging
    - codegen
    - zerolog
    - logcopter
DocType: source
Intent: long-term
Owners: []
RelatedFiles: []
ExternalSources:
    - /tmp/log-chatgpt-proposal.md
Summary: "Imported source proposal for zerolog-backed package-scoped configurable logging and code generation."
LastUpdated: 2026-05-25T10:00:00-04:00
WhatFor: "Original handoff spec that informed the LOGCOPTER-001 design guide."
WhenToUse: "Use when checking the design guide against the source proposal."
---

Below is an implementation handoff spec.

It assumes `zerolog` remains the backend. `zerolog` supports child loggers, per-logger levels via `Logger.Level`, a global process-wide clamp via `zerolog.SetGlobalLevel`, and levels including `trace`, `debug`, `info`, `warn`, `error`, `fatal`, `panic`, and `disabled`. The design below should use per-area child loggers, not the global level, for normal filtering. ([Go Packages][1])

---

# Spec: package-scoped configurable logging for Go/zerolog

## 1. Goals

Build a logging system that allows code to use a package-local logger:

```go
log.Debug().Msg("cache lookup")
log.Trace().Str("template", name).Msg("render template")
```

without manually constructing that logger in every package.

The system must support:

* `zerolog` as the logging backend.
* Per-package or per-area log levels.
* Hierarchical level inheritance.
* Config-file-driven log levels.
* Library/package prefixes, so reusable libraries can avoid naming collisions.
* Generated `logx.go` files per package.
* Runtime config reload.
* Unknown areas auto-registering for discovery.
* No-op behavior before application logging is configured.

## 2. Non-goals

This system should not:

* Replace `zerolog`.
* Parse emitted JSON to filter logs.
* Use hooks as the primary filtering mechanism.
* Require callers to write `logx.For("...")` at every call site.
* Require `go generate` during `go build`; `go generate` is explicitly run separately from builds. ([Go][2])

## 3. Terminology

### Area

An **area** is a stable logging namespace.

Examples:

```text
app.view.render
app.db.sql
app.http.server
lib.ble.rx
lib.ble.tx
```

Areas are used for config matching.

### Package prefix

A **package prefix** is a caller-defined namespace prefix added during code generation.

Examples:

```text
app
lib.ble
lib.protocol
github.com.acme.widget
```

This allows libraries to generate package loggers without colliding with application package names.

### Effective level

The **effective level** is the level that applies to a logger after hierarchical lookup.

Example config:

```yaml
logging:
  level: info
  areas:
    app.view: debug
    app.view.render: trace
```

Resolution:

```text
app.view.render.partial -> trace
app.view.assets         -> debug
app.db                  -> info
```

## 4. Config format

### 4.1 Minimal config

```yaml
logging:
  level: info
  areas:
    app.view: debug
    app.view.render: trace
    app.db: warn
    lib.ble: error
```

Meaning:

* Default level is `info`.
* `app.view` and descendants default to `debug`.
* `app.view.render` and descendants default to `trace`.
* `app.db` and descendants default to `warn`.
* `lib.ble` and descendants default to `error`.

### 4.2 Recommended full config

```yaml
logging:
  output: stderr
  format: json
  level: info
  caller: false
  timestamp: true

  areas:
    app: info
    app.view: debug
    app.view.render: trace
    app.db: warn
    app.http: info

    lib.ble: error
    lib.ble.rx: warn
    lib.protocol: debug
```

### 4.3 Supported levels

The config parser must accept:

```text
trace
debug
info
warn
warning
error
fatal
panic
off
none
disabled
```

Mapping:

```text
off / none / disabled -> zerolog.Disabled
warning               -> zerolog.WarnLevel
```

Other values should use `zerolog.ParseLevel`.

### 4.4 Invalid config behavior

If reload receives invalid config:

* Do not partially apply it.
* Keep the previous valid config.
* Return/report the error.

Example:

```text
invalid logging level for area "app.db": "verbose"
```

### 4.5 Unknown config areas

The manager should expose known/generated areas. Config entries that do not match any known area should not fail startup by default, because code may be behind build tags or plugins.

Recommended behavior:

* Warn on unknown config areas.
* Provide a strict validation mode for CI.

```yaml
logging:
  strict_areas: false
```

## 5. Area naming rules

### 5.1 Generated area formula

Generated area should be:

```text
<area-prefix>.<relative-package-path-with-slashes-replaced-by-dots>
```

Example:

```text
module path:       github.com/acme/server
package path:      github.com/acme/server/internal/view/render
strip prefix:      github.com/acme/server/internal
area prefix:       app
generated area:    app.view.render
```

For a library:

```text
module path:       github.com/acme/ble
package path:      github.com/acme/ble/rx
strip prefix:      github.com/acme/ble
area prefix:       lib.ble
generated area:    lib.ble.rx
```

### 5.2 Package prefix is mandatory for libraries

Libraries should not generate areas like:

```text
rx
tx
protocol
```

They should generate namespaced areas:

```text
lib.ble.rx
lib.ble.tx
lib.ble.protocol
```

or:

```text
github.com.acme.ble.rx
github.com.acme.ble.tx
```

### 5.3 Hierarchical matching

Use longest-prefix matching, split on `.`.

Example:

```text
lib.ble.rx.decoder.frame
```

Lookup order:

```text
lib.ble.rx.decoder.frame
lib.ble.rx.decoder
lib.ble.rx
lib.ble
lib
default
```

## 6. Runtime package: `logx`

### 6.1 Package location

For an application:

```text
internal/logx
```

For a reusable organization-wide library:

```text
github.com/acme/logx
```

Generated files import this package.

### 6.2 Public API

```go
package logx

import "github.com/rs/zerolog"

type Config struct {
	Output    string            `yaml:"output" json:"output"`
	Format    string            `yaml:"format" json:"format"`
	Level     string            `yaml:"level" json:"level"`
	Caller    bool              `yaml:"caller" json:"caller"`
	Timestamp bool              `yaml:"timestamp" json:"timestamp"`
	Areas     map[string]string `yaml:"areas" json:"areas"`
}

type Manager struct {
	// internal
}

type Logger struct {
	// small wrapper, not zerolog.Logger directly
}

func Configure(base zerolog.Logger, cfg Config) error
func DefaultManager() *Manager

func Package(area string) Logger
func For(area string) Logger

func Areas() []string
func EffectiveLevel(area string) zerolog.Level
func ValidateAreas(strict bool) []AreaWarning
```

`Package` and `For` may be aliases. Preferred naming:

```go
var log = logx.Package("app.view.render")
```

### 6.3 Logger wrapper API

The wrapper should expose normal `zerolog` event-style methods:

```go
func (l Logger) Trace() *zerolog.Event
func (l Logger) Debug() *zerolog.Event
func (l Logger) Info() *zerolog.Event
func (l Logger) Warn() *zerolog.Event
func (l Logger) Error() *zerolog.Event
func (l Logger) Fatal() *zerolog.Event
func (l Logger) Panic() *zerolog.Event
func (l Logger) WithLevel(level zerolog.Level) *zerolog.Event
func (l Logger) Raw() zerolog.Logger
```

The wrapper must resolve the current underlying `zerolog.Logger` at call time:

```go
func (l Logger) Debug() *zerolog.Event {
	return l.manager.resolve(l.area).Debug()
}
```

Do **not** make generated package loggers raw `zerolog.Logger` values. A raw logger freezes the level at initialization time.

### 6.4 Initialization behavior

Before `Configure` is called:

```go
var log = logx.Package("app.view.render")
```

must be safe.

Default behavior before configuration:

* Use `zerolog.Nop()`, or
* Use a disabled base logger.

No package should emit logs before app configuration unless the application explicitly sets an early bootstrap logger.

### 6.5 Runtime reload behavior

`Configure` may be called multiple times.

On reload:

* Parse all config first.
* If parsing succeeds, atomically replace the manager state.
* Already-created `logx.Logger` wrappers must observe new levels on future calls.
* Already-created raw `zerolog.Logger` values returned by `Raw()` are not reloadable. Document this.

### 6.6 Auto-registration

When this is called:

```go
logx.Package("lib.ble.rx")
```

the manager should add `lib.ble.rx` to a known-area set.

Unknown areas:

* Should inherit the nearest configured parent.
* Should be visible via `logx.Areas()`.
* Should not be automatically written to the config file.

### 6.7 Level filtering model

Filtering should be done by per-area child logger levels:

```go
base.With().
	Str("area", area).
	Logger().
	Level(effectiveLevel)
```

Do not set `zerolog.SetGlobalLevel(zerolog.InfoLevel)` when any area may need `debug` or `trace`. The global level is process-wide; it should only be used as an emergency clamp or left at a permissive value. `zerolog` documents `SetGlobalLevel` as affecting all loggers. ([GitHub][3])

### 6.8 Area field

Every area logger should include:

```json
"area": "app.view.render"
```

Default field name:

```text
area
```

Optional future config:

```yaml
logging:
  area_field: logger
```

Do not implement this unless needed.

## 7. Generated package files

### 7.1 Generated file content

Each package gets a generated `logx.go`:

```go
// Code generated by genlogx; DO NOT EDIT.

package render

import "github.com/acme/project/internal/logx"

var log = logx.Package("app.view.render")
```

### 7.2 Package usage

Inside the package:

```go
func Render(name string) {
	log.Trace().Str("template", name).Msg("render start")
	log.Debug().Msg("cache lookup")
}
```

No manual area string appears in normal package code.

### 7.3 Collision rules

Generated variable name defaults to:

```go
log
```

Generator must support override:

```bash
-var logger
```

Generated result:

```go
var logger = logx.Package("app.view.render")
```

This is needed if a package already uses a local `log` identifier.

## 8. Code generator: `genlogx`

### 8.1 Tool location

Recommended:

```text
tools/genlogx/main.go
```

or separate module:

```text
github.com/acme/logx/cmd/genlogx
```

### 8.2 Required flags

```text
-logx-import string
    Import path for the logx package.

-area-prefix string
    Prefix prepended to generated areas.

-strip-prefix string
    Import path prefix removed before converting package path to area.

-out string
    Output filename. Default: logx.go.

-var string
    Generated variable name. Default: log.

-include-main bool
    Whether to generate loggers for package main. Default: false.

-areas-out string
    Optional path for generated registry file containing all discovered areas.
```

### 8.3 Example: application

```bash
go run ./tools/genlogx \
  -logx-import github.com/acme/server/internal/logx \
  -strip-prefix github.com/acme/server/internal \
  -area-prefix app \
  ./internal/...
```

Package:

```text
github.com/acme/server/internal/view/render
```

Generated area:

```text
app.view.render
```

### 8.4 Example: library

```bash
go run ./tools/genlogx \
  -logx-import github.com/acme/logx \
  -strip-prefix github.com/acme/ble \
  -area-prefix lib.ble \
  ./...
```

Package:

```text
github.com/acme/ble/rx/decoder
```

Generated area:

```text
lib.ble.rx.decoder
```

### 8.5 `go generate` integration

Add a file:

```go
//go:build tools

package tools

//go:generate go run ./genlogx -logx-import github.com/acme/server/internal/logx -strip-prefix github.com/acme/server/internal -area-prefix app ../internal/...
```

Run:

```bash
go generate ./tools
```

`go generate` works by scanning source files for `//go:generate` directives and running those commands; it is not part of normal `go build`, so CI should explicitly run generation or verify generated files are current. ([Go][2])

### 8.6 Generated area registry

If `-areas-out` is set, generate:

```go
// Code generated by genlogx; DO NOT EDIT.

package logareas

var All = []string{
	"app.db",
	"app.http.server",
	"app.view.render",
	"lib.ble.rx",
}
```

Use cases:

* Config validation.
* Admin/debug UI.
* Documentation.
* Detecting stale config entries.

## 9. Application usage

### 9.1 Startup

```go
base := zerolog.New(os.Stderr).
	With().
	Timestamp().
	Logger()

err := logx.Configure(base, logx.Config{
	Output:    "stderr",
	Format:    "json",
	Level:     "info",
	Timestamp: true,
	Areas: map[string]string{
		"app.view":        "debug",
		"app.view.render": "trace",
		"app.db":          "warn",
		"lib.ble":         "error",
		"lib.ble.rx":      "warn",
	},
})
if err != nil {
	panic(err)
}
```

### 9.2 Package code

```go
package render

func Render(name string) {
	log.Trace().Str("template", name).Msg("render template")
}
```

### 9.3 Config reload

```go
func ReloadLogging(cfg logx.Config) error {
	base := zerolog.New(os.Stderr).
		With().
		Timestamp().
		Logger()

	return logx.Configure(base, cfg)
}
```

On success, existing package loggers observe the new config.

## 10. Library usage patterns

### 10.1 Internal organization libraries

For libraries controlled by the same organization, generated package loggers are acceptable:

```go
var log = logx.Package("lib.ble.rx")
```

The application controls levels:

```yaml
logging:
  level: info
  areas:
    lib.ble: error
    lib.ble.rx: trace
```

### 10.2 Public libraries

For public libraries, do not force global logging as the only option.

Recommended pattern:

```go
type Options struct {
	Logger logx.Logger
}

type Client struct {
	log logx.Logger
}

func New(opts Options) *Client {
	l := opts.Logger
	if l.IsZero() {
		l = logx.Package("lib.acme.client")
	}

	return &Client{log: l}
}
```

This allows the application to inject its own logger.

Alternative:

```go
type Options struct {
	Logger zerolog.Logger
}
```

But this loses reload behavior unless the app passes a reload-aware wrapper.

## 11. Implementation notes

### 11.1 Manager state

Use immutable state snapshots behind an atomic value or lock.

Recommended internal shape:

```go
type state struct {
	base         zerolog.Logger
	defaultLevel zerolog.Level
	configured   map[string]zerolog.Level
	known        map[string]struct{}
	cache        map[string]zerolog.Logger
}
```

Reload builds a new `state` and swaps it in.

### 11.2 Longest-prefix level resolution

```go
func resolveLevel(area string, configured map[string]zerolog.Level, fallback zerolog.Level) zerolog.Level {
	for current := area; current != ""; {
		if level, ok := configured[current]; ok {
			return level
		}

		i := strings.LastIndex(current, ".")
		if i < 0 {
			break
		}

		current = current[:i]
	}

	return fallback
}
```

### 11.3 Event allocation and performance

This wrapper adds one manager lookup per log call.

Acceptable for normal use. If an extremely hot path needs maximum performance, allow:

```go
raw := log.Raw()
raw.Debug().Msg("hot path")
```

Document that `Raw()` is not reload-aware after it is captured.

### 11.4 Caller information

If the application enables `zerolog.Caller()`, wrapper methods may affect reported call sites.

Preferred design:

```go
log.Debug().Msg("x")
```

where only `Debug()` is wrapped and the event chain remains zerolog-native.

Avoid helper methods like:

```go
log.DebugMsg("x")
```

because they are more likely to hide the real caller.

## 12. Testing requirements

### 12.1 Unit tests

Test:

* Default level.
* Exact area match.
* Parent prefix match.
* Unknown area inherits default.
* `disabled` drops logs.
* Reload changes effective levels.
* Generated logger created before `Configure` works after `Configure`.
* Invalid reload does not replace previous config.
* `Areas()` includes generated/registered areas.
* Config warning for stale area.

### 12.2 Generator tests

Test:

* Application package path to area conversion.
* Library package path to area conversion.
* `-area-prefix`.
* `-strip-prefix`.
* `-var`.
* Skipping `main` by default.
* Generated code is `gofmt`-formatted.
* Generated code compiles.

## 13. Acceptance criteria

Implementation is complete when:

* Packages can use generated `var log = logx.Package("...")`.
* App config can set `trace` for one area and `warn` for another.
* Library prefixes work cleanly, for example `lib.ble.rx`.
* Existing generated loggers observe config reload.
* Unknown areas auto-register and inherit sensible defaults.
* `go generate` can regenerate all `logx.go` files.
* CI can verify generated files are up to date.
* No package needs to manually create a zerolog child logger just to get area filtering.

## 14. Recommended defaults

Use these defaults unless there is a reason not to:

```yaml
logging:
  output: stderr
  format: json
  level: info
  caller: false
  timestamp: true
  areas: {}
```

Generator defaults:

```text
-out logx.go
-var log
-include-main false
```

Application area prefix:

```text
app
```

Library area prefix:

```text
lib.<library-name>
```

Example final config:

```yaml
logging:
  level: info
  areas:
    app.view.render: trace
    app.db: warn
    app.http: info

    lib.ble: error
    lib.ble.rx: trace
    lib.protocol: debug
```

[1]: https://pkg.go.dev/github.com/rs/zerolog?utm_source=chatgpt.com "zerolog package - github.com/rs/zerolog"
[2]: https://go.dev/blog/generate?utm_source=chatgpt.com "Generating code"
[3]: https://github.com/rs/zerolog?utm_source=chatgpt.com "rs/zerolog: Zero Allocation JSON Logger"

