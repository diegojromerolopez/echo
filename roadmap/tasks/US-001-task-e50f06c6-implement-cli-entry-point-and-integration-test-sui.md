# Task: Implement CLI entry point and integration test suite

- **ID**: `task-e50f06c6`
- **Story ID**: `US-001`
- **Status**: `PENDING`
- **Change Type**: `FEATURE`
- **Depends On**: `Scaffold Go module and implement core Echoer package with unit tests`
- **Target Files**: `cmd/echo/main.go`, `tests/cli_integration_test.go`

## Description

1. Create `cmd/echo/main.go` as the CLI composition root. It passes `os.Args[1:]` to `echoer.Echo` and writes the returned string followed by exactly one `\n` to `os.Stdout`. Stderr remains empty and process exits with code `0`.
2. Create `tests/cli_integration_test.go` which uses `go build` to compile `cmd/echo` into a temporary binary inside `t.TempDir()`. Run subprocess tests executing Edge Case Matrix rows 1-5 (`echo-cli hello world`, `echo-cli hello`, `echo-cli`, `echo-cli "hello world" foo`, `echo-cli ""`), asserting exact stdout (including `\n`), zero-byte stderr, and exit code `0`. Ensure files remain under 500 lines.
