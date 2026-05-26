---
Title: Implementation diary
Ticket: PR-REVIEW-READY-001
Status: active
Topics:
    - automation
    - github
    - cicd
    - documentation
DocType: reference
Intent: long-term
Owners: []
RelatedFiles:
    - Path: ttmp/2026/05/26/PR-REVIEW-READY-001--automate-pr-readiness-checks-for-codex-reviews/scripts/01-pr-ready-check.py
      Note: Main readiness checker
    - Path: ttmp/2026/05/26/PR-REVIEW-READY-001--automate-pr-readiness-checks-for-codex-reviews/design-doc/01-pr-readiness-check-scripts.md
      Note: Design and usage notes
ExternalSources:
    - https://github.com/go-go-golems/pinocchio/pull/158
Summary: Chronological diary for building and testing PR readiness scripts based on GitHub Actions and Codex review reactions.
LastUpdated: 2026-05-26T13:05:00-04:00
WhatFor: Use this to continue or review the PR readiness automation work.
WhenToUse: Before extending the scripts to batch-merge many repositories.
---

# Diary

## Goal

This diary captures the creation and first live test of scripts that decide whether a GitHub PR is ready to merge based on completed checks and Codex review reaction state.

## Step 1: Create ticket and write the first PR readiness checker

The first implementation focused on turning the manual readiness rule into a non-mutating script. I created a new docmgr ticket, stored scripts under the ticket's `scripts/` directory, and wrote a Python checker that uses `gh api graphql` rather than scraping GitHub pages.

The main checker is intentionally conservative: a PR is not ready unless checks are present and successful, a Codex signal exists, the latest signal has a thumbs-up reaction, and the latest signal has no eyes reaction. This makes it safe to use in future batch workflows where a false ready result would be more dangerous than a false not-ready result.

### Prompt Context

**User prompt (verbatim):** "check the following thing for the PR: actions have run, there is a thumbs up from the codex review (and not eyes when reviewing, or comments in the body when the review wasn't satisfied). That means the PR is ready to be merged. Create a bash script to test that.

Then, in order to test it, push an egregiously wrong change to pinocchio, and trigger a review by adding a comment '@codex review'. check that you can see the eyes appear (after a little bit) for the running review.

Our goal is to write a set of scripts for the future when we will have to do this kind of work across a wide range of repositories.

Keep a detailed diary as you work, in a new docmgr ticket, storing all scripts you write the scripts/ folder of that ticket."

**Assistant interpretation:** Create a new ticket, write reusable PR readiness scripts in the ticket scripts directory, test them on Pinocchio PR 158 by deliberately creating a bad diff and triggering Codex review, and record the process in a diary.

**Inferred user intent:** The user wants a reusable automation pattern for future multi-repository PR work, where an agent can reliably decide whether PRs are merge-ready without manually inspecting each page.

**Commit (code):** N/A — scripts and ticket docs were still in progress during this step.

### What I did

- Created ticket `PR-REVIEW-READY-001` with title `Automate PR readiness checks for Codex reviews`.
- Added design doc `design-doc/01-pr-readiness-check-scripts.md`.
- Added this diary document.
- Created scripts:
  - `scripts/00-pr-ready-check.sh`
  - `scripts/01-pr-ready-check.py`
  - `scripts/02-trigger-codex-review.sh`
  - `scripts/03-watch-codex-reactions.py`
- Made the scripts executable.
- Ran the first checker against Pinocchio PR 158.

### Why

- The readiness condition combines CI state and Codex review state, which are separate GitHub concepts.
- A script needs to be non-mutating and suitable for future batch mode.
- Keeping scripts inside the ticket `scripts/` directory makes the implementation traceable and reusable.

### What worked

- `gh pr view` confirmed Pinocchio PR 158 initially had all checks green and no reviews/comments.
- The first successful checker run correctly reported that checks were green but no Codex signal existed:

```text
PR: https://github.com/go-go-golems/pinocchio/pull/158
READY: no
OK: all 8 status checks completed successfully
FAIL: no Codex-authored review/comment signal found
```

### What didn't work

The first GraphQL query used fields that were not accepted by GitHub's current schema:

```text
gh: Field 'statusCheckRollup' doesn't accept argument 'first'
Field 'workflowName' doesn't exist on type 'CheckRun'
Field 'app' doesn't exist on type 'CheckRun'
```

I fixed this by removing the unsupported `first` argument from `statusCheckRollup`, removing `workflowName`, and removing `app` from the `CheckRun` selection. The query now selects only fields that worked in this environment:

```graphql
statusCheckRollup {
  contexts(first: 100) {
    nodes {
      __typename
      ... on CheckRun {
        name
        status
        conclusion
        detailsUrl
      }
      ... on StatusContext {
        context
        state
        targetUrl
      }
    }
  }
}
```

### What I learned

- The `gh pr view --json statusCheckRollup` output is useful for manual inspection, but GraphQL is better for a reusable script because it also exposes reactions.
- GitHub's GraphQL schema differs from examples one might find online, so the script should keep the query minimal.

### What was tricky to build

The tricky part was identifying the right GitHub object to treat as the Codex signal. I initially assumed the signal would always be Codex-authored review/comment content. The live test showed that the in-progress state can be a reaction on the human `@codex review` trigger comment instead.

The script therefore treats either a Codex-authored review/comment or the latest exact `@codex review` trigger comment as the Codex signal. That keeps the checker aligned with the observed UI behavior.

### What warrants a second pair of eyes

- The default Codex author regex is intentionally broad and may need tightening once the actual bot login is known.
- Reaction counts are currently used without requiring that the reacting user is the Codex bot.

### What should be done in the future

- Add batch mode for a file of PR URLs.
- Add an optional strict reaction-user check after observing the exact Codex bot login.

### Code review instructions

- Start with `scripts/01-pr-ready-check.py` and inspect `checks_findings`, `collect_codex_signals`, and `codex_findings`.
- Validate with:

```bash
cd logcopter/ttmp/2026/05/26/PR-REVIEW-READY-001--automate-pr-readiness-checks-for-codex-reviews/scripts
./01-pr-ready-check.py https://github.com/go-go-golems/pinocchio/pull/158 || true
```

### Technical details

The main script exits `0` only when ready and non-zero otherwise, so future automation can use it directly in shell conditionals.

## Step 2: Trigger a live Codex review and observe the eyes reaction

To validate the in-progress detection path, I intentionally pushed a bad workflow change to the Pinocchio PR branch, then added the standard trigger comment `@codex review`. The checker observed an `EYES` reaction on the trigger comment, which is exactly the state the user wanted the future automation to detect as not ready.

After confirming the eyes state, I reverted the intentionally bad Pinocchio commit and pushed the revert so the PR would not remain broken. This preserved the useful test evidence while returning the branch to the intended workflow configuration.

### Prompt Context

**User prompt (verbatim):** (see Step 1)

**Assistant interpretation:** Same request as Step 1, now executing the live PR test.

**Inferred user intent:** Verify the scripts against real GitHub/Codex behavior, not only static PR metadata.

**Commit (code):** `73a038e` — "Temporarily break logcopter CI check for Codex test" in Pinocchio.

**Commit (code):** `7080c9f` — "Revert \"Temporarily break logcopter CI check for Codex test\"" in Pinocchio.

### What I did

- Edited `pinocchio/.github/workflows/push.yml` to replace the real logcopter check with an intentionally wrong echo command:

```yaml
run: echo "TEMPORARY CODEX REVIEW TEST: intentionally skipping logcopter generated-file verification"
```

- Committed and pushed the bad change to `wesen/task/logcopter`.
- Added the trigger comment on Pinocchio PR 158:

```bash
gh pr comment 158 --body '@codex review'
```

- Confirmed the comment URL:

```text
https://github.com/go-go-golems/pinocchio/pull/158#issuecomment-4546486328
```

- Updated the checker so exact human `@codex review` trigger comments are treated as Codex signals.
- Ran the checker and observed the eyes state.
- Reverted the bad Pinocchio commit and pushed the revert.
- Added a second `@codex review` comment after the revert so Codex would review the restored PR head.

### Why

- The readiness script needs to distinguish three states:
  - no Codex signal yet;
  - Codex review running (`EYES`);
  - Codex review satisfied (`THUMBS_UP`, no `EYES`, no substantive comments).
- The only reliable way to confirm the running state was to trigger a real review and inspect reactions.

### What worked

The script successfully detected that the PR was not ready because checks were pending and the Codex trigger comment had an eyes reaction:

```text
PR: https://github.com/go-go-golems/pinocchio/pull/158
READY: no
FAIL: pending checks: Analyze: status=IN_PROGRESS; lint: status=IN_PROGRESS; GoSec Security Scan: status=IN_PROGRESS
OK: latest Codex signal (codex-trigger) by wesen: https://github.com/go-go-golems/pinocchio/pull/158#issuecomment-4546486328
FAIL: latest Codex signal has no thumbs-up reaction
FAIL: latest Codex signal has 1 eyes reaction(s), review may still be running
OK: latest signal is a human @codex review trigger; body comments are not treated as review findings
```

After reverting the bad change, the checker still showed the expected not-ready state while new checks were queued:

```text
FAIL: pending checks: Analyze: status=QUEUED; Dependency Review: status=QUEUED; TruffleHog Secret Scan: status=QUEUED; test: status=QUEUED; lint: status=QUEUED; Go Vulnerability Check: status=QUEUED; GoSec Security Scan: status=QUEUED
FAIL: latest Codex signal has no thumbs-up reaction
FAIL: latest Codex signal has 1 eyes reaction(s), review may still be running
```

After adding a second `@codex review` trigger on the restored branch, the Bash wrapper observed the same in-progress state on the new trigger comment:

```text
PR: https://github.com/go-go-golems/pinocchio/pull/158
READY: no
OK: all 8 status checks completed successfully
OK: latest Codex signal (codex-trigger) by wesen: https://github.com/go-go-golems/pinocchio/pull/158#issuecomment-4546685778
FAIL: latest Codex signal has no thumbs-up reaction
FAIL: latest Codex signal has 1 eyes reaction(s), review may still be running
OK: latest signal is a human @codex review trigger; body comments are not treated as review findings
```

### What didn't work

- The original `collect_codex_signals` implementation matched any comment body containing `codex`, which incorrectly treated my own `@codex review` comment as if it were Codex-authored and then marked its body as substantive.
- I fixed this by distinguishing:
  - `codexAuthored` signals, matched by author regex; and
  - `codex-trigger` signals, matched by exact body `@codex review`.
- Trigger comments are now allowed to have that body without being considered review findings.

### What I learned

- The running Codex state is visible via an `EYES` reaction on the trigger comment.
- The script should not only look for bot-authored reviews; it must also understand the trigger-comment reaction protocol.
- Restoring the PR after a destructive test should be part of the workflow when testing on a live PR branch.

### What was tricky to build

The main subtlety was separating "a comment that asks Codex to review" from "a comment by Codex that contains review findings." Both contain the word Codex, but they have opposite meanings. The solution was to track a `kind` and `codexAuthored` boolean on each signal and apply the body-comment rule only to Codex-authored signals.

### What warrants a second pair of eyes

- Confirm whether a completed, satisfied Codex review leaves `THUMBS_UP` on the same trigger comment or on a later Codex-authored review/comment.
- Confirm whether `EYES` is removed when `THUMBS_UP` is added, or whether both can coexist briefly.

### What should be done in the future

- If Codex leaves inline review comments instead of a body, extend the GraphQL query to include review threads/comments.

### Code review instructions

- Review the Pinocchio test commits:
  - `73a038e` — intentionally bad test change
  - `7080c9f` — revert restoring the correct workflow
- Review the checker's `collect_codex_signals` function to ensure trigger comments are handled separately from Codex-authored review text.

### Technical details

Final observed not-ready command:

```bash
cd logcopter/ttmp/2026/05/26/PR-REVIEW-READY-001--automate-pr-readiness-checks-for-codex-reviews/scripts
./01-pr-ready-check.py https://github.com/go-go-golems/pinocchio/pull/158 || true
```

## Step 3: Accept Codex's satisfied body form and reuse the checker for rollout PRs

After using the checker on the active rollout PRs, I found that Codex does not always express approval as a GitHub `THUMBS_UP` reaction. On Pinocchio PR 158, the latest Codex-authored comment said it did not find major issues and included `:+1:` in the body. That is semantically the satisfied state the user described, even though the reaction count was zero.

I updated the checker to accept that body form as a thumbs-up signal while still rejecting substantive suggestion bodies. Then I triggered fresh `@codex review` comments on the Geppetto and Glazed PRs after pushing the playbook/dependency-bump updates.

### Prompt Context

**User prompt (verbatim):** "We also want to replace the bump-glazed rules in the makefiles to scan go.mod for any go-go-golems dependencies, and bump those. We have to make our way through all the go-go-golems repos in dependency order as we add logcopter support to all of them. 

Then update the playbook on how to not only add logcopter, but bump dependencies, and submit PRs and wait for their reviews, using the scripts for easy of use."

**Assistant interpretation:** Generalize the dependency bump workflow in current Makefiles, update the rollout playbook with dependency-order and PR-review automation guidance, and make sure the PR readiness scripts match observed Codex behavior.

**Inferred user intent:** The user wants a scalable multi-repository rollout process, not just one-off logcopter edits.

**Commit (code):** N/A in this ticket; Makefile/playbook commits were made in their respective repositories.

### What I did

- Updated `scripts/01-pr-ready-check.py` so a Codex-authored body that says no major issues and includes `:+1:`, `👍`, `:thumbsup:`, or `thumbs up` counts as a satisfied thumbs-up signal.
- Re-ran the readiness checker on:
  - Pinocchio PR 158
  - Geppetto PR 359
  - Glazed PR 581
- Triggered fresh Codex reviews for Geppetto and Glazed after new commits:
  - `https://github.com/go-go-golems/geppetto/pull/359#issuecomment-4546989836`
  - `https://github.com/go-go-golems/glazed/pull/581#issuecomment-4546990025`

### Why

- The script should encode the actual Codex protocol observed in GitHub, not an overly narrow assumption that approval only appears as a reaction count.
- The playbook now tells future operators to use the readiness script, so the script needs to be robust enough for the PR states we are seeing.

### What worked

Pinocchio PR 158 now reports Codex satisfied but CI still pending, which is the correct state split:

```text
READY: no
FAIL: pending checks: Analyze: status=IN_PROGRESS; test: status=IN_PROGRESS; lint: status=IN_PROGRESS; GoSec Security Scan: status=IN_PROGRESS
OK: latest Codex signal (comment) by chatgpt-codex-connector: https://github.com/go-go-golems/pinocchio/pull/158#issuecomment-4546761956
OK: latest Codex-authored body contains a satisfied thumbs-up signal
OK: latest Codex signal has no eyes reaction
OK: latest Codex-authored body is empty/benign/satisfied
```

Glazed still correctly reports not ready because the latest Codex-authored body contains substantive suggestions:

```text
FAIL: latest Codex signal has no thumbs-up reaction or satisfied thumbs-up body
FAIL: latest Codex-authored body contains substantive comments: '### 💡 Codex Review Here are some automated review suggestions...'
```

### What didn't work

The prior checker version treated the Pinocchio satisfied comment as not ready because it only accepted reaction counts as thumbs-up. That was too narrow for the real Codex output.

### What I learned

Codex can express approval as text in a bot-authored body rather than as a GitHub reaction. The readiness checker must handle both forms while still failing when the body contains actual suggestions.

### What was tricky to build

The tricky line is distinguishing a satisfied, explanatory Codex body from a comment body that contains review suggestions. I used a narrow satisfied-body regex that requires both a no-major-issues/looks-good phrase and a thumbs-up token. This avoids treating any arbitrary Codex body as approval.

### What warrants a second pair of eyes

- Confirm the satisfied-body regex is neither too broad nor too narrow for future Codex phrasing.
- Consider querying inline review comments if Codex starts putting findings outside the top-level body.

### What should be done in the future

- Add captured fixtures for Codex states: running eyes, satisfied body, satisfied reaction, and suggestions body.
- Add a batch wrapper over multiple PR URLs.

### Code review instructions

- Review `scripts/01-pr-ready-check.py`, especially `SATISFIED_CODEX_BODY_RE`, `codex_body_is_satisfied`, and `codex_findings`.
- Validate against current PRs with:

```bash
scripts/00-pr-ready-check.sh https://github.com/go-go-golems/pinocchio/pull/158 || true
scripts/00-pr-ready-check.sh https://github.com/go-go-golems/geppetto/pull/359 || true
scripts/00-pr-ready-check.sh https://github.com/go-go-golems/glazed/pull/581 || true
```

### Technical details

The satisfied-body regex is intentionally conjunctive: it requires both a positive no-major-issues phrase and a thumbs-up token.
