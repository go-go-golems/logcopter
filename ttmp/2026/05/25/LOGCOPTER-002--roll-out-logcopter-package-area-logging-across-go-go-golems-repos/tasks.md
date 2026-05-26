# Tasks

## Phase 0 — Ticket setup and analysis

- [x] Create LOGCOPTER-002 ticket workspace.
- [x] Add primary design/implementation guide document.
- [x] Add investigation diary document.
- [x] Inspect current Glazed logcopter integration and dependency state.
- [x] Inspect Pinocchio logging initialization entry points.
- [x] Inspect representative Geppetto logging-heavy packages.
- [x] Inspect Clay logging wrapper/deprecation state.
- [x] Count current `github.com/rs/zerolog/log` usage by repository.
- [x] Decide rollout starts with Glazed dependency cleanup and self-adoption.
- [x] Define canonical area prefix convention: `go-go-golems.<repo>.<path-tail>`.

## Phase 1 — Glazed first implementation pass

- [x] Update `glazed/go.mod` to require `github.com/go-go-golems/logcopter v0.0.1`.
- [x] Remove local `replace github.com/go-go-golems/logcopter => ../logcopter` from Glazed.
- [x] Register `logcopter-gen` using the Go tool mechanism so `go tool logcopter-gen` works.
- [x] Add a repository-local `go:generate` entry point for Glazed package logger generation.
- [x] Ensure `go generate ./...` regenerates Glazed logcopter package logger files.
- [x] Run `go mod tidy` in Glazed.
- [x] Generate Glazed package loggers for `./pkg/...` with prefix `go-go-golems.glazed`.
- [x] Remove or alias `github.com/rs/zerolog/log` imports in converted Glazed packages.
- [x] Add generated-file check for Glazed package loggers.
- [x] Run targeted Glazed tests: `go test ./pkg/cmds/fields ./pkg/cmds/logging ./cmd/glaze`.
- [x] Run targeted tests for converted Glazed packages.
- [x] Smoke test `go run ./cmd/glaze --log-area go-go-golems.glazed.pkg.help=debug help logging-section-reference`.
- [x] Commit Glazed changes (`69733764289f9939cb0cbccad71b76b7466c59d8`).

## Phase 2 — Pinocchio rollout

- [x] Update Pinocchio to a Glazed version containing published logcopter integration through the workspace checkout while keeping module dependency compatible with current release state.
- [x] Add direct `github.com/go-go-golems/logcopter v0.0.1` dependency.
- [x] Generate package loggers with prefix `go-go-golems.pinocchio`.
- [x] Convert package diagnostics away from global zerolog imports where generated loggers exist.
- [x] Replace touched `clay.InitGlazed` calls with direct Glazed logging setup.
- [x] Validate root `pinocchio` command bootstrap.
- [x] Validate `web-chat` and `simple-chat-agent` command bootstrap.
- [x] Commit Pinocchio changes (`c2161fc`).

## Phase 3 — Geppetto rollout

- [x] Add direct `github.com/go-go-golems/logcopter v0.0.1` dependency.
- [x] Generate package loggers with prefix `go-go-golems.geppetto` for `./pkg/...`.
- [x] Convert high-value AI provider packages first.
- [x] Convert inference tool/middleware/event packages.
- [x] Preserve explicit `zerolog.Logger` injection APIs.
- [x] Validate targeted Geppetto package tests.
- [x] Commit Geppetto changes (`998f3651`).

## Phase 4 — Clay rollout

- [x] Add direct `github.com/go-go-golems/logcopter v0.0.1` dependency.
- [x] Generate package loggers with prefix `go-go-golems.clay` for `./pkg/...`.
- [x] Convert watcher, repositories, filters, SQL, and workerpool diagnostics.
- [x] Keep Clay logging initialization helpers deprecated.
- [x] Validate `go test ./pkg/...`.
- [x] Commit Clay changes (`440a77d`).

## Phase 5 — Cross-repository validation and docs

- [x] Write a Glazed help tutorial playbook for converting go-go-golems packages to logcopter.
- [x] Update rollout playbook with correct non-mutating CI drift-check ordering.
- [x] Replace hand-maintained `bump-glazed` targets with generic `bump-go-go-golems` targets.
- [x] Update rollout playbook with dependency-order, dependency-bump, PR submission, and Codex readiness guidance.
- [x] Align Geppetto, Pinocchio, Clay, and go-template CI with the playbook ordering.
- [x] Validate the playbook is discoverable through `glaze help`.
- [x] Use the playbook to guide the Geppetto rollout.
- [x] Use the playbook to guide the Pinocchio rollout.
- [x] Smoke test CLI log configuration with explicit logcopter config files.
- [ ] Verify one `go-go-golems.glazed...` area can be verbose while default remains quiet.
- [ ] Verify one `go-go-golems.pinocchio...` area can be configured through Pinocchio CLI.
- [ ] Verify one `go-go-golems.geppetto...` area can be configured through a Pinocchio command path.
- [x] Verify Clay package diagnostics/config smoke through a Glazed-configured example application.
- [ ] Update repository READMEs or help docs with cross-repository area prefix examples.
- [x] Update LOGCOPTER-002 diary and changelog after each implementation phase.
- [ ] Run `docmgr doctor --ticket LOGCOPTER-002 --stale-after 30`.
- [ ] Upload final bundle to reMarkable.
