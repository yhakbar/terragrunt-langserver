# Notes

- Replace Hermit with mise.
- Lint. The formatting is hard to read.
- Logging goes to stderr, no option to log to file.
- Quite a lot of dead code.
- Testing is lacking.
- Have to preserve dead code to implement protocol.Server. Not a good trade-off.
- DidChange not implemented. Active buffers won't go through the LS.
- Godoc missing from packages.
- In `lsp/document/workspace.go`, why not use `sync.Map`?
- In `lsp/protocol`, why are `tsclient.go`, `uri.go` and `tsserver.go` changed?
  - It seems largely related to getting packages replaced, but is this scalable?
  - Would prefer static copy of vendored packages, or use of Go modules.
  - It's also vendoring content from a project's `internal` directory,
    which goes against what the `internal` directory is for.

