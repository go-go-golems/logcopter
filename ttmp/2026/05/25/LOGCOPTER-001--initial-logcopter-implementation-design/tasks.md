# Tasks

## Completed

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

## TODO / follow-up for implementation

- [ ] Rename module from `github.com/go-go-golems/XXX` to `github.com/go-go-golems/logcopter`.
- [ ] Scaffold `pkg/logcopter` runtime package.
- [ ] Implement level parsing and hierarchical area lookup.
- [ ] Implement reload-aware manager and logger wrapper.
- [ ] Implement `cmd/logcopter-gen` generator.
- [ ] Implement Glazed adapter package.
- [ ] Add runtime, generator, and Glazed integration tests.
- [ ] Add examples and README documentation.
- [ ] Verify Glazed config-file support for `areas` map vs `log-area` string-list syntax.
