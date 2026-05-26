# Changelog

## 2026-05-25

- Initial workspace created


## 2026-05-25

Created cross-repository logcopter rollout analysis and implementation guide

### Related Files

- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-002--roll-out-logcopter-package-area-logging-across-go-go-golems-repos/design-doc/01-cross-repository-logcopter-rollout-analysis-and-implementation-guide.md — Primary intern-facing analysis and rollout guide
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-002--roll-out-logcopter-package-area-logging-across-go-go-golems-repos/reference/01-investigation-diary.md — Investigation diary step 1
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-002--roll-out-logcopter-package-area-logging-across-go-go-golems-repos/tasks.md — Phased task checklist


## 2026-05-25

Validated LOGCOPTER-002 and uploaded analysis bundle to reMarkable

### Related Files

- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-002--roll-out-logcopter-package-area-logging-across-go-go-golems-repos/reference/01-investigation-diary.md — Recorded validation and reMarkable upload
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/vocabulary.yaml — Added clay


## 2026-05-25

Started Glazed transition and added go tool/go generate tasks

### Related Files

- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-002--roll-out-logcopter-package-area-logging-across-go-go-golems-repos/reference/01-investigation-diary.md — Recorded Step 3 start of Glazed transition
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-002--roll-out-logcopter-package-area-logging-across-go-go-golems-repos/tasks.md — Added Go tool and go generate tasks to Phase 1


## 2026-05-25

Completed Glazed package logger transition (commit 69733764289f9939cb0cbccad71b76b7466c59d8)

### Related Files

- /home/manuel/workspaces/2026-05-25/logcopter/glazed/.github/workflows/push.yml — Added generated package logger freshness check
- /home/manuel/workspaces/2026-05-25/logcopter/glazed/Makefile — Added logcopter generation/check targets
- /home/manuel/workspaces/2026-05-25/logcopter/glazed/go.mod — Requires published logcopter v0.0.1 and registers logcopter-gen tool
- /home/manuel/workspaces/2026-05-25/logcopter/glazed/logcopter_generate.go — Repository-local go generate entry point for logcopter package loggers
- /home/manuel/workspaces/2026-05-25/logcopter/glazed/pkg/cmds/logging/init.go — Aliases global zerolog package as zlog while package diagnostics use generated log
- /home/manuel/workspaces/2026-05-25/logcopter/glazed/pkg/helpers/files/temp-files_test.go — Aliases standard library log as stdlog to avoid generated log collision
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-002--roll-out-logcopter-package-area-logging-across-go-go-golems-repos/reference/01-investigation-diary.md — Recorded Step 4 implementation diary
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-002--roll-out-logcopter-package-area-logging-across-go-go-golems-repos/tasks.md — Marked Glazed Phase 1 tasks complete


## 2026-05-25

Uploaded updated LOGCOPTER-002 bundle after Glazed transition

### Related Files

- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-002--roll-out-logcopter-package-area-logging-across-go-go-golems-repos/reference/01-investigation-diary.md — Recorded Step 5 reMarkable upload


## 2026-05-25

Added Glazed logcopter rollout playbook and ported Geppetto/Pinocchio with config smoke tests

### Related Files

- /home/manuel/workspaces/2026-05-25/logcopter/geppetto/logcopter_generate.go — Geppetto go generate entry point for logcopter package loggers
- /home/manuel/workspaces/2026-05-25/logcopter/glazed/pkg/doc/tutorials/logcopter-package-rollout-playbook.md — New Glazed tutorial playbook for logcopter migration
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-002--roll-out-logcopter-package-area-logging-across-go-go-golems-repos/reference/01-investigation-diary.md — Recorded Steps 6-9 detailed implementation diary
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-002--roll-out-logcopter-package-area-logging-across-go-go-golems-repos/tasks.md — Marked playbook
- /home/manuel/workspaces/2026-05-25/logcopter/pinocchio/logcopter_generate.go — Pinocchio go generate entry point for package and command subpackage loggers


## 2026-05-25

Uploaded updated LOGCOPTER-002 bundle after Geppetto and Pinocchio rollout

### Related Files

- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-002--roll-out-logcopter-package-area-logging-across-go-go-golems-repos/reference/01-investigation-diary.md — Recorded Step 10 reMarkable upload


## 2026-05-25

Embedded logcopter docs for Glazed import and completed Clay rollout

### Related Files

- /home/manuel/workspaces/2026-05-25/logcopter/clay/examples/simple/logging_layer_example.go — Example now uses Glazed logging setup directly
- /home/manuel/workspaces/2026-05-25/logcopter/clay/logcopter_generate.go — Clay go generate entry point for logcopter package loggers
- /home/manuel/workspaces/2026-05-25/logcopter/glazed/cmd/glaze/main.go — Loads logcopter help docs by importing the embedded FS
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/pkg/doc/doc.go — Exports embedded logcopter help docs as an fs.FS
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-002--roll-out-logcopter-package-area-logging-across-go-go-golems-repos/reference/01-investigation-diary.md — Recorded Steps 11-12 implementation diary
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/25/LOGCOPTER-002--roll-out-logcopter-package-area-logging-across-go-go-golems-repos/tasks.md — Marked Clay rollout complete

