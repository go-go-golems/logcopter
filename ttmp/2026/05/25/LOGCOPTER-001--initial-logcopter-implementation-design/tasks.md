# Tasks

## Completed research and preparation

- [x] Create LOGCOPTER-001 ticket workspace.
- [x] Add primary design document.
- [x] Add investigation diary.
- [x] Import `/tmp/log-chatgpt-proposal.md` into `sources/`.
- [x] Read and summarize the imported proposal.
- [x] Inspect current logcopter skeleton.
- [x] Study Glazed logging section and initialization code.
- [x] Study Pinocchio's early and final logging bootstrap.
- [x] Write intern-facing implementation guide with APIs, pseudocode, diagrams, and file references.
- [x] Update ticket index and diary.
- [x] Preliminary dependency cleanup: deprecate Clay logging/config initialization helpers and document direct Glazed replacements.
- [x] Revise the design so logcopter is a utility package and Glazed's existing logging package owns command/config integration.
- [x] Specify `fields.TypeKeyValue` as the Glazed field type for CLI area overrides.

## Phase 0 — Repository rename and baseline scaffold

- [x] Rename module from `github.com/go-go-golems/XXX` to `github.com/go-go-golems/logcopter`.
- [x] Rename `cmd/XXX` to `cmd/logcopter-gen`.
- [x] Replace template README with a short logcopter README.
- [x] Decide minimum Go/toolchain version for the module.
- [x] Add runtime dependency on `github.com/rs/zerolog`.
- [x] Add generator dependency on `golang.org/x/tools/go/packages`.
- [x] Add test dependency such as `github.com/stretchr/testify` if used. (Not needed yet; Phase 1 uses the standard library `testing` package.)
- [x] Run `go mod tidy`.
- [x] Run `go test ./...` to establish a clean scaffold baseline.

## Phase 1 — Runtime level and area primitives

- [x] Create `pkg/logcopter/level.go`.
- [x] Implement `ParseLevel` with aliases: `warning`, `off`, `none`, `disabled`.
- [x] Add table tests for valid levels.
- [x] Add table tests for invalid levels with clear errors.
- [x] Create `pkg/logcopter/areas.go`.
- [x] Implement area normalization/validation helper.
- [x] Implement longest-prefix level resolution.
- [x] Test exact match, parent match, default fallback, and invalid area strings.

## Phase 2 — Reload-aware runtime manager

- [x] Create `pkg/logcopter/config.go` with `Config` and `AreaWarning`.
- [x] Create `pkg/logcopter/manager.go` with immutable state snapshots.
- [x] Create default manager in `pkg/logcopter/global.go`.
- [x] Implement `Configure(base zerolog.Logger, cfg Config) error`.
- [x] Ensure invalid reload does not replace the previous valid state.
- [x] Implement `Package(area string)` / `For(area string)` and known-area registration.
- [x] Implement `Areas()`, `EffectiveLevel(area)`, and `ValidateAreas(strict)`.
- [x] Test pre-configuration no-op behavior.
- [x] Test reload changes future log calls on existing wrappers.
- [x] Test unknown configured areas in warning and strict modes.

## Phase 3 — Logger wrapper API

- [x] Create `pkg/logcopter/logger.go`.
- [x] Implement `Trace`, `Debug`, `Info`, `Warn`, `Error`, `Fatal`, `Panic`, and `WithLevel`.
- [x] Implement `Raw() zerolog.Logger` with documentation that captured raw loggers are not reload-aware.
- [x] Implement `Area()` and `IsZero()` helpers.
- [x] Add tests that emitted events include the `area` field.
- [x] Add tests that disabled levels drop logs.
- [x] Add caller-output test if `Caller()` support is enabled through the base logger.
- [x] Benchmark manager lookup overhead before adding any cache.

## Phase 4 — Output construction helpers

- [x] Create `pkg/logcopter/output.go` if output construction belongs in logcopter runtime.
- [x] Support stderr default output.
- [x] Support stdout where needed.
- [x] Support JSON output.
- [x] Support text console output.
- [x] Decide whether rotating file output stays in Glazed only or is shared with logcopter. Decision: rotating file output stays in Glazed; logcopter only provides small stream/format helpers.
- [x] Test output writer selection separately from area filtering.

## Phase 5 — Code generator

- [x] Implement `cmd/logcopter-gen/main.go` flag parsing.
- [x] Implement package discovery with `go/packages`.
- [x] Implement `-logcopter-import`.
- [x] Implement `-strip-prefix`.
- [x] Implement `-area-prefix`.
- [x] Implement `-out` with default `logcopter.go`.
- [x] Implement `-var` with default `log`.
- [x] Implement `-include-main` defaulting to false.
- [x] Implement `-areas-out` registry generation.
- [x] Implement `-dry-run`.
- [x] Implement `-check` for CI.
- [x] Render generated files with an explicit import alias: `import logcopter "..."`.
- [x] Run generated output through `gofmt`.
- [x] Test package path to area conversion.
- [x] Test strip-prefix errors.
- [x] Test generated source compilation in a temp module.

## Phase 6 — Glazed key-value field support for area overrides

- [x] Inspect current `fields.TypeKeyValue` parser behavior in `glazed/pkg/cmds/fields`.
- [x] Add support for both `key:value` and `key=value` inputs if not already supported.
- [x] Keep existing colon syntax backward-compatible.
- [x] Add parser tests for repeated values: `app.view:debug`, `app.db=warn`.
- [x] Add parser tests for comma-separated Cobra `StringSlice` values if supported by pflag.
- [x] Add clear parse errors for malformed entries such as `app.view` or `app.view:`.
- [x] Confirm decoded values can populate `map[string]string` in a settings struct.

## Phase 6.5 — Logcopter help entries surfaced through Glaze

- [x] Add Markdown-only logcopter help entries under `logcopter/pkg/doc` without importing Glazed.
- [x] Write a general architecture guide help entry.
- [x] Write a tutorial help entry for generated package loggers and Glazed logging configuration.
- [x] Update `glazed/cmd/glaze` to discover and load logcopter Markdown help entries when the logcopter checkout is present.
- [x] Validate with `go run ./cmd/glaze help logcopter-logging-architecture` from the Glazed checkout.
- [x] Decide whether release builds should copy/embed logcopter help docs into Glazed, or keep this as workspace-only discovery. Decision for initial version: keep workspace-only discovery and revisit before release packaging.

## Phase 7 — In-place Glazed logging integration

- [x] Update `glazed/pkg/cmds/logging/section.go`.
- [x] Add `log-config` as a repeatable string-list field for explicit logcopter-only config/profile files.
- [x] Add `log-area` as a `fields.TypeKeyValue` field.
- [x] Add or preserve canonical config support for `logging.areas` map.
- [x] Add `strict-log-areas` boolean field.
- [x] Update `LoggingSettings` with `LogConfigFiles []string`, `LogAreas map[string]string`, `Areas map[string]string`, and `StrictAreas bool`.
- [x] Update `AddLoggingSectionToRootCommand` with manual persistent flags for `--log-config`, `--log-area`, and `--strict-log-areas`.
- [x] Update `InitLoggerFromCobra` to read `--log-config` and `--log-area` through the same code paths as section parsing.
- [x] Update `SetupLoggingFromValues` / `GetLoggingSettings` to normalize section values.
- [x] Implement deterministic merge order: defaults, app logging section, explicit logcopter config files in order, direct CLI flags.
- [x] Support both `logging:`-wrapped and direct logcopter-only profile file shapes.
- [x] Update `InitLoggerFromSettings` to configure logcopter's default manager.
- [x] Stop using `zerolog.SetGlobalLevel` as the normal filtering mechanism when logcopter is active.
- [x] Update `InitEarlyLoggingFromArgs` to preserve and parse explicit log config and area override flags before command discovery.
- [x] Add Glazed unit tests for CLI flags, parsed values, explicit log config files, config map values, early logging, and strict area validation.

## Phase 8 — Cross-repository validation

- [x] Run `go test ./...` in `logcopter`.
- [x] Run targeted Glazed logging tests.
- [x] Run broader Glazed tests if feasible. Completed by Glazed commit hook (`go test ./...`, lint, gosec, govulncheck).
- [x] Run at least a Pinocchio command bootstrap smoke test.
- [x] Verify existing Glazed applications can keep their current imports and setup calls.
- [x] Verify one area can emit `trace` while another remains at `warn`.
- [x] Verify a generated package logger created before configuration observes later config.

## Phase 9 — Documentation and examples

- [x] Update logcopter README with runtime and generator examples.
- [x] Add `examples/basic` for native runtime use.
- [x] Add `examples/library-prefix` for reusable package prefixes.
- [x] Add `examples/glazed-cli` or a Glazed docs example showing in-place Glazed logging setup.
- [x] Update `glazed/pkg/doc/topics/logging-section.md` with area override examples.
- [x] Document CLI examples: `--log-area app.view:debug`, `--log-area app.view=debug`, and `--log-config ~/.config/logcopter/profiles/dev.yaml`.
- [x] Document YAML examples using application `logging.areas`.
- [x] Document standalone logcopter profile files with both `logging:`-wrapped and direct shapes.
- [x] Document `Raw()` reload caveat.
- [x] Document global zerolog level interaction.

## Phase 10 — Release readiness

- [x] Add CI check for generated files using `logcopter-gen -check`.
- [x] Add `go generate` instructions.
- [x] Add changelog entry for logcopter initial release.
- [x] Add migration note for Glazed logging users.
- [x] Tag or prepare release once runtime, generator, and Glazed integration are validated. Prepared release notes and generated-file CI; no tag was created pending review.
