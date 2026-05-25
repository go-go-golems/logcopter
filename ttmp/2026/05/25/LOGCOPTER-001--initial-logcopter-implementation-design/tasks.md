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
- [ ] Add generator dependency on `golang.org/x/tools/go/packages`.
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
- [ ] Add caller-output test if `Caller()` support is enabled through the base logger.
- [ ] Benchmark manager lookup overhead before adding any cache.

## Phase 4 — Output construction helpers

- [ ] Create `pkg/logcopter/output.go` if output construction belongs in logcopter runtime.
- [ ] Support stderr default output.
- [ ] Support stdout where needed.
- [ ] Support JSON output.
- [ ] Support text console output.
- [ ] Decide whether rotating file output stays in Glazed only or is shared with logcopter.
- [ ] Test output writer selection separately from area filtering.

## Phase 5 — Code generator

- [ ] Implement `cmd/logcopter-gen/main.go` flag parsing.
- [ ] Implement package discovery with `go/packages`.
- [ ] Implement `-logcopter-import`.
- [ ] Implement `-strip-prefix`.
- [ ] Implement `-area-prefix`.
- [ ] Implement `-out` with default `logcopter.go`.
- [ ] Implement `-var` with default `log`.
- [ ] Implement `-include-main` defaulting to false.
- [ ] Implement `-areas-out` registry generation.
- [ ] Implement `-dry-run`.
- [ ] Implement `-check` for CI.
- [ ] Render generated files with an explicit import alias: `import logcopter "..."`.
- [ ] Run generated output through `gofmt`.
- [ ] Test package path to area conversion.
- [ ] Test strip-prefix errors.
- [ ] Test generated source compilation in a temp module.

## Phase 6 — Glazed key-value field support for area overrides

- [ ] Inspect current `fields.TypeKeyValue` parser behavior in `glazed/pkg/cmds/fields`.
- [ ] Add support for both `key:value` and `key=value` inputs if not already supported.
- [ ] Keep existing colon syntax backward-compatible.
- [ ] Add parser tests for repeated values: `app.view:debug`, `app.db=warn`.
- [ ] Add parser tests for comma-separated Cobra `StringSlice` values if supported by pflag.
- [ ] Add clear parse errors for malformed entries such as `app.view` or `app.view:`.
- [ ] Confirm decoded values can populate `map[string]string` in a settings struct.

## Phase 7 — In-place Glazed logging integration

- [ ] Update `glazed/pkg/cmds/logging/section.go`.
- [ ] Add `log-config` as a repeatable string-list field for explicit logcopter-only config/profile files.
- [ ] Add `log-area` as a `fields.TypeKeyValue` field.
- [ ] Add or preserve canonical config support for `logging.areas` map.
- [ ] Add `strict-log-areas` boolean field.
- [ ] Update `LoggingSettings` with `LogConfigFiles []string`, `LogAreas map[string]string`, `Areas map[string]string`, and `StrictAreas bool`.
- [ ] Update `AddLoggingSectionToRootCommand` with manual persistent flags for `--log-config`, `--log-area`, and `--strict-log-areas`.
- [ ] Update `InitLoggerFromCobra` to read `--log-config` and `--log-area` through the same code paths as section parsing.
- [ ] Update `SetupLoggingFromValues` / `GetLoggingSettings` to normalize section values.
- [ ] Implement deterministic merge order: defaults, app logging section, explicit logcopter config files in order, direct CLI flags.
- [ ] Support both `logging:`-wrapped and direct logcopter-only profile file shapes.
- [ ] Update `InitLoggerFromSettings` to configure logcopter's default manager.
- [ ] Stop using `zerolog.SetGlobalLevel` as the normal filtering mechanism when logcopter is active.
- [ ] Update `InitEarlyLoggingFromArgs` to preserve and parse explicit log config and area override flags before command discovery.
- [ ] Add Glazed unit tests for CLI flags, parsed values, explicit log config files, config map values, early logging, and strict area validation.

## Phase 8 — Cross-repository validation

- [ ] Run `go test ./...` in `logcopter`.
- [ ] Run targeted Glazed logging tests.
- [ ] Run broader Glazed tests if feasible.
- [ ] Run at least a Pinocchio command bootstrap smoke test.
- [ ] Verify existing Glazed applications can keep their current imports and setup calls.
- [ ] Verify one area can emit `trace` while another remains at `warn`.
- [ ] Verify a generated package logger created before configuration observes later config.

## Phase 9 — Documentation and examples

- [ ] Update logcopter README with runtime and generator examples.
- [ ] Add `examples/basic` for native runtime use.
- [ ] Add `examples/library-prefix` for reusable package prefixes.
- [ ] Add `examples/glazed-cli` or a Glazed docs example showing in-place Glazed logging setup.
- [ ] Update `glazed/pkg/doc/topics/logging-section.md` with area override examples.
- [ ] Document CLI examples: `--log-area app.view:debug`, `--log-area app.view=debug`, and `--log-config ~/.config/logcopter/profiles/dev.yaml`.
- [ ] Document YAML examples using application `logging.areas`.
- [ ] Document standalone logcopter profile files with both `logging:`-wrapped and direct shapes.
- [ ] Document `Raw()` reload caveat.
- [ ] Document global zerolog level interaction.

## Phase 10 — Release readiness

- [ ] Add CI check for generated files using `logcopter-gen -check`.
- [ ] Add `go generate` instructions.
- [ ] Add changelog entry for logcopter initial release.
- [ ] Add migration note for Glazed logging users.
- [ ] Tag or prepare release once runtime, generator, and Glazed integration are validated.
