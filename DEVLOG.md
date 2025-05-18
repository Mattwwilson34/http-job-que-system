# Development Log: Go Job Queue

This is an ongoing technical journal documenting my progress, design decisions,
experiments, and reflections while building an HTTP-based job queue system in
Go.

---

## 📅 [2025-05-18]

### What I worked on:
- Implemented a custom logger for the application that writes logs to a file
(`log.txt`) in the root directory

### Problems or blockers:
- None in the traditional sense, but the main challenge was figuring out how to
implement file-based logging in Go using idiomatic patterns

### Decisions made and why:
- Opted to roll my own file-based logger using Go’s standard `log` and `os`
packages
- Decided to fail fast and exit the application (`log.Fatal`) if logger
initialization fails, to ensure visibility into startup issues and guarantee
that logging is available from the beginning
- Chose to write logs to a persistent file so I can track application flow,
debug issues more easily, and retain logs between sessions

### What I learned:
- How to use Go's `os.OpenFile` with flags like `os.O_CREATE`, `os.O_WRONLY`,
and `os.O_APPEND` to control how the log file is opened
- How Go uses octal notation (`0644`) for file permissions, and how to interpret
those permission bits
- The difference between `os.Executable()` and `os.Getwd()`:
  - `os.Executable()` points to the temporary binary path created by `go run .`
  - `os.Getwd()` returns the actual working directory, which is what I needed to
  place my log file at the expected location
- Basic architectural thinking around startup order, initialization, and the
importance of early failure in observability-critical subsystems like logging

### Next steps:
- Learn how to interact with my http server via my go application

## 📅 [2025-05-17]

### What I worked on:
- Plan an initialize project repo
- Add pre-commit hooks to prevent push to main
- Add PR template
- Add first implementation of http server

### Problems or blockers:
- none

### Decisions made and why:
- Prevent pushes to main to force me to create PR for all work

### What I learned:
- How to start a basic http server in Go 

### Next steps:
- Setup a custom logger that outputs to a system file

