# Tasks

## Phase 1 — Ticket setup

- [x] Create `PR-REVIEW-READY-001` ticket workspace.
- [x] Add design document.
- [x] Add implementation diary.
- [x] Store scripts in the ticket `scripts/` directory.

## Phase 2 — Script implementation

- [x] Write Bash entry point for the PR readiness checker.
- [x] Write non-mutating PR readiness checker implementation.
- [x] Query GitHub status checks via GraphQL.
- [x] Query PR comments/reviews and reaction groups via GraphQL.
- [x] Detect successful/pending/failing Actions checks.
- [x] Detect `@codex review` trigger comments.
- [x] Detect `EYES` as an in-progress Codex review signal.
- [x] Detect missing `THUMBS_UP` as not ready.
- [x] Add helper script to trigger Codex review.
- [x] Add helper script to poll for Codex reactions.

## Phase 3 — Live Pinocchio test

- [x] Push intentionally wrong Pinocchio workflow change.
- [x] Trigger Codex with `@codex review`.
- [x] Observe `EYES` reaction with the checker.
- [x] Revert intentionally wrong Pinocchio change.
- [x] Trigger Codex again on the restored Pinocchio branch and observe `EYES` on the new trigger comment.
- [ ] Observe final completed `THUMBS_UP` state after Codex finishes.

## Phase 4 — Hardening/future work

- [ ] Add batch mode over a file of PR URLs.
- [ ] Restrict reaction checks to the actual Codex bot login once confirmed.
- [ ] Query inline review comments/review threads if Codex uses them for unsatisfied reviews.
- [ ] Add shell wrapper that summarizes many PRs in a compact table.
- [ ] Decide whether the script should optionally merge ready PRs or remain check-only.
