# Changelog

## 2026-05-25

- Created LOGCOPTER-001 ticket workspace for the initial logcopter implementation design.
- Imported `/tmp/log-chatgpt-proposal.md` into `sources/01-log-chatgpt-proposal.md`.
- Studied Glazed logging support in `glazed/pkg/cmds/logging/*`, including section definition, Cobra flag wiring, logger initialization, and early logging parsing.
- Studied `pinocchio/cmd/pinocchio/main.go` as a complex bootstrap example with early logging before dynamic command loading.
- Wrote `design-doc/01-initial-logcopter-implementation-guide.md`, including architecture, APIs, pseudocode, diagrams, implementation phases, tests, risks, and file references.
- Wrote `reference/01-investigation-diary.md` following the ticket diary format.
- Updated `index.md` and `tasks.md` for handoff.

## 2026-05-25

Completed research/design deliverable for initial logcopter implementation guide

### Related Files

- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-001--initial-logcopter-implementation-design/design-doc/01-initial-logcopter-implementation-guide.md — Primary intern-facing implementation guide
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-001--initial-logcopter-implementation-design/reference/01-investigation-diary.md — Chronological investigation diary
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-001--initial-logcopter-implementation-design/sources/01-log-chatgpt-proposal.md — Imported source proposal


## 2026-05-25

Preliminary dependency cleanup: deprecated Clay logging/config init helpers in favor of direct Glazed setup

### Related Files

- /home/manuel/workspaces/2026-05-25/logcopter/clay/README.md — Updated configuration section with Glazed replacement example
- /home/manuel/workspaces/2026-05-25/logcopter/clay/pkg/init.go — Deprecated InitGlazed and Viper init helpers with porting guidance
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-001--initial-logcopter-implementation-design/design-doc/01-initial-logcopter-implementation-guide.md — Added preliminary implementation phase
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-001--initial-logcopter-implementation-design/reference/01-investigation-diary.md — Recorded Step 5 diary entry


## 2026-05-25

Updated design: logcopter remains a utility package while Glazed's existing logging section/init gain logcopter configuration support

### Related Files

- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-001--initial-logcopter-implementation-design/design-doc/01-initial-logcopter-implementation-guide.md — Replaced separate Glazed adapter design with in-place Glazed logging package changes
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-001--initial-logcopter-implementation-design/reference/01-investigation-diary.md — Recorded Step 6 design redirection
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-001--initial-logcopter-implementation-design/tasks.md — Updated implementation tasks for Glazed logging package changes


## 2026-05-25

Specified Glazed TypeKeyValue area overrides and expanded tasks into granular implementation phases

### Related Files

- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-001--initial-logcopter-implementation-design/design-doc/01-initial-logcopter-implementation-guide.md — Documents TypeKeyValue log-area syntax and parser expectations
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-001--initial-logcopter-implementation-design/reference/01-investigation-diary.md — Recorded Step 7 documentation update
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-001--initial-logcopter-implementation-design/tasks.md — Granular phase-by-phase implementation checklist


## 2026-05-25

Added explicit logcopter-only config/profile file support to the design

### Related Files

- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-001--initial-logcopter-implementation-design/design-doc/01-initial-logcopter-implementation-guide.md — Documents --log-config and profile merge order
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-001--initial-logcopter-implementation-design/reference/01-investigation-diary.md — Records Step 8 before implementation begins
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-001--initial-logcopter-implementation-design/tasks.md — Adds granular tasks for explicit logcopter profile files


## 2026-05-25

Phase 0 scaffold: renamed module and replaced placeholder command with logcopter-gen scaffold

### Related Files

- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/README.md — Initial project README
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/cmd/logcopter-gen/main.go — Generator command scaffold with planned flags
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/go.mod — Module renamed to github.com/go-go-golems/logcopter
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-001--initial-logcopter-implementation-design/reference/01-investigation-diary.md — Recorded Step 9 scaffold diary
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-001--initial-logcopter-implementation-design/tasks.md — Phase 0 checklist progress


## 2026-05-25

Phase 1 runtime primitives: implemented level parsing and area resolution helpers

### Related Files

- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/pkg/logcopter/areas.go — Area normalization and longest-prefix level resolution
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/pkg/logcopter/areas_test.go — Area validation and resolution tests
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/pkg/logcopter/level.go — ParseLevel with aliases for warning/off/none/disabled
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/pkg/logcopter/level_test.go — Level parsing table tests
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-001--initial-logcopter-implementation-design/reference/01-investigation-diary.md — Recorded Step 10 runtime primitives diary
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-001--initial-logcopter-implementation-design/tasks.md — Marked Phase 1 tasks complete


## 2026-05-25

Phase 2/3 runtime: implemented reload-aware manager and logger wrapper

### Related Files

- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/pkg/logcopter/config.go — Runtime Config and AreaWarning types
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/pkg/logcopter/global.go — Default manager package helpers
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/pkg/logcopter/logger.go — Reload-aware zerolog-style Logger wrapper
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/pkg/logcopter/manager.go — Manager state snapshots
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/pkg/logcopter/manager_test.go — Manager reload
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-001--initial-logcopter-implementation-design/reference/01-investigation-diary.md — Recorded Step 11 manager/logger diary
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-001--initial-logcopter-implementation-design/tasks.md — Marked Phase 2 and Phase 3 progress

