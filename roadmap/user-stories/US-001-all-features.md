# User Story 001: Implement Echo CLI (Full Vertical Slice)

As a user,
I want a command-line tool that prints the first argument I pass to it (or a blank line if no arguments are given),
So that I can verify basic CLI input/output behavior reliably and predictably.

---
depends_on: []
change_type: new
target_files:
  - go.mod
  - cmd/echo/main.go
  - pkg/echoer/echoer.go
  - pkg/echoer/echoer_test.go
  - tests/cli_integration_test.go
  - Makefile
---

## Overview

This is the sole, comprehensive user story for the Echo CLI project (CU < 25, small-spec rule applies). It delivers a complete, runnable, end-to-end vertical slice: a compiled binary `echo-cli` that echoes its first CLI argument to stdout, backed by a pure, testable core function and full unit + integration test coverage. This story is self-contained and requires no subsequent stories to be demonstrable.

## Requirements

### Task 1: Project Scaffold, Module & Core Echoer Logic (with Unit Tests)
- Create `go.mod` declaring module path `github.com/noctifab/echo` and Go version `1.22` or later.
- Create `pkg/echoer/echoer.go` implementing:
  ```go
  package echoer

  // Echo returns the first element of args, or "" if args is empty.
  func Echo(args []string) string
  ```
  - `Echo` MUST be a pure function with no I/O side effects, no global/package-level mutable state, and no dependency on `os.Args`.
  - Behavior:
    - `Echo([]string{"hello", "world"})` returns `"hello"`.
    - `Echo([]string{"only-one"})` returns `"only-one"`.
    - `Echo([]string{})` returns `""`.
    - `Echo(nil)` returns `""`.
    - `Echo([]string{""})` (first arg is an empty string) returns `""`.
    - `Echo([]string{"hello world", "foo"})` (first arg contains internal whitespace) returns `"hello world"` unmodified (no trimming, no splitting on internal spaces).
- Create `pkg/echoer/echoer_test.go` in the SAME task, co-located with the implementation, covering all bullet cases above using table-driven tests.
- File size constraint: `echoer.go` and `echoer_test.go` must each remain strictly under 500 lines.

### Task 2: CLI Composition Root & Binary Entry Point (with Integration Tests)
- Create `cmd/echo/main.go` as the composition root:
  - Reads `os.Args[1:]` and passes it explicitly (dependency injection, no global state) to `echoer.Echo`.
  - Writes the returned string followed by exactly one `\n` character to `os.Stdout` using buffered or direct writes (e.g., `fmt.Println` or `fmt.Fprintln(os.Stdout, ...)`).
  - Writes nothing to `os.Stderr` under normal operation.
  - Exits with status code `0` in all cases described by this spec (there are no error paths defined for this CLI).
  - The compiled binary name is `echo-cli` (built from package `github.com/noctifab/echo/cmd/echo`).
- Create `tests/cli_integration_test.go` in the SAME task, co-located conceptually with the CLI entry point delivery:
  - Uses `go build` (e.g., via `os/exec` and `testing.T.TempDir()`) to compile the `github.com/noctifab/echo/cmd/echo` package into a temporary binary path.
  - Invokes the compiled binary via subprocess execution (`exec.Command`) with various argument sets.
  - Asserts exact stdout content (byte-for-byte, including trailing newline), asserts stderr is empty, and asserts exit code is `0` for each scenario in the Edge Case Matrix below.
- File size constraint: `main.go` and `cli_integration_test.go` must each remain strictly under 500 lines.

### Task 3: Build & Quality Automation (Makefile)
- Create `Makefile` with at least two targets, functionally verified as part of this task:
  - `test`: runs `go test -v ./...` and must complete with exit code 0 when all tests pass.
  - `lint`: runs `go vet ./...` and must complete with exit code 0 with zero warnings/errors.
- Validate both targets execute successfully against the codebase produced in Tasks 1 and 2.

## Architectural & Performance Constraints
- SOLID separation: CLI argument parsing and I/O reside exclusively in `cmd/echo/main.go`; string/business logic resides exclusively in `pkg/echoer`.
- Dependency Injection: command-line arguments are passed explicitly as function parameters into `echoer.Echo`; no reliance on package-level or global state.
- No Go source file may exceed 500 lines of code.
- All code must be formatted per `go fmt` standard (no diffs when `gofmt -l .` is run).
- All code must pass `go vet ./...` with zero warnings or errors.

## Definition of Done (DoD)

1. **Public Interface & Entry Point Contracts:**
   - Public Go function: `func Echo(args []string) string` in package `github.com/noctifab/echo/pkg/echoer`.
   - Public binary executable: compiled from `github.com/noctifab/echo/cmd/echo`, producing an executable conventionally named `echo-cli` (or an equivalent temp-built binary in tests), invoked as `echo-cli [arg1] [arg2] ...`.

2. **Standard I/O & Output Formatting Invariants:**
   - `echo-cli <arg1> <arg2> ...` prints `<arg1>` followed by exactly one `\n` to stdout, then exits with code `0`.
   - `echo-cli` (no arguments) prints exactly one `\n` to stdout, then exits with code `0`.
   - Stderr MUST be empty (zero bytes) in all defined scenarios.
   - No error prefixes, warnings, or additional output of any kind are permitted; there are no defined error/failure paths for this CLI (all invocations succeed with exit code `0`).

3. **Explicit Data Representation & Edge Cases:**
   - Arguments are treated as opaque strings; no numeric parsing, trimming, or type coercion occurs.
   - Empty-string arguments (e.g., `echo-cli ""`) are treated as valid first arguments and printed as an empty line (i.e., just `\n`, since the empty string plus newline is indistinguishable from the no-args case at the output level, but must be verified via `Echo` returning `""`).
   - Arguments containing internal whitespace (e.g., `"hello world"`) are printed verbatim, unmodified.
   - No time, date, TTL, or clock-dependent behavior exists in this feature; the Deterministic Time & Mock Clock Invariant is not applicable.

4. **Comprehensive Scenario & Edge Case Matrix:**
   | # | Invocation | Expected stdout | Expected stderr | Expected exit code |
   |---|---|---|---|---|
   | 1 | `echo-cli hello world` | `hello\n` | (empty) | 0 |
   | 2 | `echo-cli hello` | `hello\n` | (empty) | 0 |
   | 3 | `echo-cli` | `\n` | (empty) | 0 |
   | 4 | `echo-cli "hello world" foo` | `hello world\n` | (empty) | 0 |
   | 5 | `echo-cli ""` | `\n` | (empty) | 0 |
   | 6 | `Echo([]string{})` (unit) | returns `""` | n/a | n/a |
   | 7 | `Echo(nil)` (unit) | returns `""` | n/a | n/a |
   | 8 | `Echo([]string{"only-one"})` (unit) | returns `"only-one"` | n/a | n/a |

5. **Verification Criteria:**
   - `go test -v ./...` MUST pass with 100% success rate (zero failing tests) and cover all rows of the Edge Case Matrix.
   - `go vet ./...` MUST report zero warnings or errors.
   - `gofmt -l .` MUST report zero files needing formatting changes.
   - `make test` and `make lint` MUST both succeed (exit code 0).

6. **Deterministic Time & Mock Clock Invariants:** Not applicable — this feature has no time, timer, TTL, or wall-clock dependency.

7. **Pinned Directory & Required File Paths:** The following exact file paths MUST all exist in the repository upon completion:
   - `go.mod`
   - `cmd/echo/main.go`
   - `pkg/echoer/echoer.go`
   - `pkg/echoer/echoer_test.go`
   - `tests/cli_integration_test.go`
   - `Makefile`

8. **Machine-Readable Story Contract:** See fenced `noctifab-contract` block below.

9. **Zero Host Package Installations Mandate:** All verification (unit tests, integration tests, `go vet`, `go fmt` checks, and Makefile targets) MUST run using only the standard Go toolchain already present in the build environment; no host OS package installations, external daemons, databases, or containerized services are required for this feature, as it has no external dependencies.

## Validation Criteria

### Unit Tests (`pkg/echoer/echoer_test.go`)
- Table-driven tests covering all `Echo` cases in the Edge Case Matrix (rows 6-8) plus the whitespace-preservation and empty-string-argument cases described above.

### Integration Tests (`tests/cli_integration_test.go`)
- Compile `github.com/noctifab/echo/cmd/echo` to a temporary binary path using `go build` invoked programmatically.
- Execute the compiled binary as a subprocess with the argument sets in Edge Case Matrix rows 1-5, asserting exact stdout bytes, empty stderr, and exit code 0 for each.

### Test Execution Command
- All tests must pass via: `go test -v ./...`

```noctifab-contract
{
  "story_id": "US-001",
  "public_contracts": [
    {
      "id": "cli.echo-with-args",
      "interface": "CLI ./cmd/echo",
      "applicable_path_prefixes": ["cmd/echo/", "pkg/echoer/", "tests/"],
      "allowed_executables": ["./cmd/echo/echo-cli", "echo-cli"],
      "exit_codes": [0],
      "stdout_contains": ["hello\n"],
      "stderr_prefixes": []
    },
    {
      "id": "cli.echo-no-args",
      "interface": "CLI ./cmd/echo",
      "applicable_path_prefixes": ["cmd/echo/", "pkg/echoer/", "tests/"],
      "allowed_executables": ["./cmd/echo/echo-cli", "echo-cli"],
      "exit_codes": [0],
      "stdout_contains": ["\n"],
      "stderr_prefixes": []
    }
  ]
}
```
