# Task: Scaffold Go module and implement core Echoer package with unit tests

- **ID**: `task-d341bc24`
- **Story ID**: `US-001`
- **Status**: `PENDING`
- **Change Type**: `FEATURE`
- **Depends On**: `[]` 
- **Target Files**: `go.mod`, `pkg/echoer/echoer.go`, `pkg/echoer/echoer_test.go`

## Description

1. Create `go.mod` declaring module path `github.com/noctifab/echo` and Go version `1.22`.
2. Create `pkg/echoer/echoer.go` implementing pure function `Echo(args []string) string`. It returns the first element of `args`, or `""` if `args` is empty or nil. It must perform no I/O, rely on no global state, and preserve internal whitespace unmodified.
3. Create `pkg/echoer/echoer_test.go` with table-driven unit tests covering all required cases: multiple args (`["hello", "world"]`), single arg (`["only-one"]`), empty slice (`[]string{}`), nil slice (`nil`), empty string first arg (`[]string{""}`), and whitespace preservation (`["hello world", "foo"]`). Keep both `.go` files under 500 lines.
