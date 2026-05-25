---
Title: Initial logcopter implementation design
Ticket: LOGCOPTER-001
Status: active
Topics:
    - logging
    - codegen
    - zerolog
    - glazed
    - logcopter
DocType: index
Intent: long-term
Owners: []
RelatedFiles: []
ExternalSources:
    - sources/01-log-chatgpt-proposal.md
Summary: "Ticket workspace for the initial logcopter design and implementation guide."
LastUpdated: 2026-05-25T09:55:00-04:00
WhatFor: "Track the design and future implementation of logcopter's zerolog-backed area-scoped logging runtime, generator, and Glazed adapter."
WhenToUse: "Use before implementing or reviewing the initial logcopter release."
---

# Initial logcopter implementation design

## Overview

LOGCOPTER-001 captures the first implementation plan for logcopter: a configurable Go logging system built on `zerolog`, with generated package-local loggers and hierarchical area-scoped log levels.

The key design goal is that application packages can use a generated local logger:

```go
var log = logcopter.Package("app.view.render")
```

and application config can set different levels for different package/area boundaries:

```yaml
logging:
  level: info
  areas:
    app.view.render: trace
    app.db: warn
```

## Key links

- [Initial logcopter implementation guide](design-doc/01-initial-logcopter-implementation-guide.md)
- [Investigation diary](reference/01-investigation-diary.md)
- [Imported source proposal](sources/01-log-chatgpt-proposal.md)
- [Tasks](tasks.md)
- [Changelog](changelog.md)

## Current status

Status: **active**.

The design/research deliverable is complete. Implementation has not started. The next phase is to scaffold the runtime package and generator according to the guide.

## Important conclusions

- Logcopter should keep `zerolog` as the backend.
- Generated package loggers must be reload-aware wrappers, not raw `zerolog.Logger` values.
- Normal filtering should use per-area child logger levels, not `zerolog.SetGlobalLevel`.
- The Glazed adapter should mirror Glazed's section and Cobra initialization UX while avoiding Glazed's current global-level filtering path.
- Pinocchio demonstrates why an early logging initializer is useful before dynamic command discovery.

## Topics

- logging
- codegen
- zerolog
- glazed
- logcopter

## Structure

- `design-doc/` — primary design and implementation guide.
- `reference/` — investigation diary.
- `sources/` — imported source proposal.
- `tasks.md` — completed research tasks and implementation follow-ups.
- `changelog.md` — ticket changes and decisions.
