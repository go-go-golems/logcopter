# Changelog

## 2026-05-26

- Initial workspace created


## 2026-05-26

Created PR readiness scripts and verified Codex eyes reaction on Pinocchio PR 158

### Related Files

- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/26/PR-REVIEW-READY-001--automate-pr-readiness-checks-for-codex-reviews/scripts/01-pr-ready-check.py — Main non-mutating readiness checker for checks and Codex reaction state
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/26/PR-REVIEW-READY-001--automate-pr-readiness-checks-for-codex-reviews/scripts/02-trigger-codex-review.sh — Helper to post @codex review
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/26/PR-REVIEW-READY-001--automate-pr-readiness-checks-for-codex-reviews/scripts/03-watch-codex-reactions.py — Helper to poll for Codex reaction signals
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/26/PR-REVIEW-READY-001--automate-pr-readiness-checks-for-codex-reviews/design-doc/01-pr-readiness-check-scripts.md — Design and usage reference
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/26/PR-REVIEW-READY-001--automate-pr-readiness-checks-for-codex-reviews/reference/01-implementation-diary.md — Detailed implementation diary
- /home/manuel/workspaces/2026-05-25/logcopter/pinocchio/.github/workflows/push.yml — Temporarily modified for Codex test and then restored by revert

## 2026-05-26

Added Bash wrapper and retriggered Codex on restored Pinocchio branch

### Related Files

- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/26/PR-REVIEW-READY-001--automate-pr-readiness-checks-for-codex-reviews/scripts/00-pr-ready-check.sh — Shell entry point for future operators and batch scripts
- /home/manuel/workspaces/2026-05-25/logcopter/logcopter/ttmp/2026/05/26/PR-REVIEW-READY-001--automate-pr-readiness-checks-for-codex-reviews/reference/01-implementation-diary.md — Records second Codex trigger on restored branch
