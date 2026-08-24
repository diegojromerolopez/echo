# Echo CLI (`echo-cli`)

[![Noctifab Generated](https://img.shields.io/badge/Generated%20by-Noctifab-blue)](https://github.com/diegojromerolopez/noctifab)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A minimal, deterministic command-line interface (CLI) utility written in Go that outputs the first argument it receives to standard output.

This project is an autonomous software factory project designed, specified, implemented, and verified end-to-end by **[Noctifab](https://github.com/diegojromerolopez/noctifab)** — the autonomous dark factory agent platform for software engineering.

---

## 1. Overview

`echo-cli` is a lightweight Go utility adhering to SOLID design principles and pure domain boundaries:
- **With Arguments**: `echo-cli <arg1> <arg2> ...` prints the first argument (`<arg1>`) followed by a newline `\n` to `stdout`, exiting with code `0`.
- **Without Arguments**: `echo-cli` prints a single newline `\n` to `stdout`, exiting with code `0`.
- **Pure Domain Boundary**: Argument processing resides in `pkg/echoer`, completely decoupled from CLI argument parsing and I/O in `cmd/echo`.

---

## 2. Project Architecture

```
echo/
├── SPEC.md                       # Technical specification & definition of done
├── go.mod                        # Go module definition (Go 1.22+)
├── Makefile                      # Project test & quality automation
├── cmd/
│   └── echo/
│       └── main.go               # Composition root (CLI entrypoint)
├── pkg/
│   └── echoer/
│       ├── echoer.go             # Pure domain logic (Echo function)
│       └── echoer_test.go        # Table-driven unit tests
└── tests/
    └── cli_integration_test.go   # Subprocess CLI integration tests
```

---

## 3. Quick Start

### Building and Running
```bash
# Build the binary
go build -o bin/echo-cli cmd/echo/main.go

# Run with arguments
./bin/echo-cli hello world
# Output: hello

# Run without arguments
./bin/echo-cli
# Output: (blank line)
```

### Running Tests & Linting
```bash
# Run all unit and integration tests
make test

# Run static analysis
make lint
```

---

## 4. Built with Noctifab

This repository was generated autonomously by **[Noctifab](https://github.com/diegojromerolopez/noctifab)** via its multi-agent orchestration pipeline:
- **Product Manager Agent**: Decomposed `SPEC.md` into actionable user stories with explicit Definition of Done (DoD) criteria and black-box contract scenarios.
- **Planner Agent**: Synthesized an optimal dependency graph of atomic development tasks.
- **Generator & Tester Agents**: Executed test-driven development (TDD) in isolated worktree sandboxes.
- **Sovereign Pre-Merge Gate**: Verified 100% test pass consensus before integrating into the release trunk.

For more information on the dark factory architecture, visit the [Noctifab repository on GitHub](https://github.com/diegojromerolopez/noctifab).
