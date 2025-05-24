# Development Log: Go Job Queue

This is an ongoing technical journal documenting my progress, design decisions,
experiments, and reflections while building an HTTP-based job queue system in
Go.

---

# Development Log: Go Job Queue
This is an ongoing technical journal documenting my progress, design decisions,
experiments, and reflections while building an HTTP-based job queue system in
Go.

---
## 📅 [2025-05-24]

### What I worked on:
- Fixed HTTP handler flow control bug where POST requests fell through to 403 responses
- Replaced server-crashing `log.Fatal` calls with proper HTTP error responses
- Learned Go's string vs byte slice differences for HTTP response writing
- Created structured JSON response types for consistent API responses
- Debugged HTTP response header/body ordering issues

### Problems or blockers:
- Missing `return` statements in POST handlers caused execution to continue to default 403
- `ResponseWriter.Write()` expects `[]byte`, not `string` - required type conversion
- Tried to capture JSON data from `json.NewEncoder().Encode()` which returns error, not data
- Set headers after body was already written (wrong HTTP order)

### Decisions made and why:
- Added explicit `return` statements after all response paths - Go requires explicit control flow
- Used `http.StatusBadRequest` for JSON parse errors instead of crashing server
- Chose `fmt.Fprint()` over `ResponseWriter.Write()` to avoid manual string-to-bytes conversion

### What I learned:
- Go's explicit control flow - blocks don't auto-return from functions
- HTTP response order: headers → status → body (strict protocol requirement)
- `json.NewEncoder(w).Encode()` writes directly to response, `json.Marshal()` returns bytes
- String/byte slice distinction reflects network communication happening in bytes

### Next steps:
- Add proper Content-Type headers for JSON responses
- Implement input validation for Job struct
- Begin actual job queue storage logic

## 📅 [2025-05-23]

### What I worked on:

* Refined understanding of Go's `net/http` package, including both server-side request handling and client-side testing.
* Explored differences between `http.Client{}` and `&http.Client{}`—why the pointer form is preferred for managing internal state like connection pools.
* Built and tested HTTP route handlers with method restrictions.
* Created client-side test code using `http.NewRequest()` and `client.Do()` to allow testing of all HTTP methods, beyond just `GET` and `POST`.
* Debugged test failures related to `DELETE` requests, focusing on the interaction between client tests and server error handling.
* Improved server log messaging and evaluated proper status codes and error responses.

### Problems or blockers:

* Tests for `DELETE` requests failed due to the server crashing (`log.Fatal`) when decoding an empty JSON body—preventing status code validation.
* Misinterpreted test failure as a `t.Errorf` for unexpected status codes, when it was actually a `t.Fatalf` caused by a transport error (due to server termination).
* Initial confusion around method constants and the separation between `http.NewRequest()` vs convenience methods (`http.Get()`, etc.).

### Decisions made and why:

* **HTTP Client Instantiation**: Use `&http.Client{}` to enable reusability and proper management of internal state, especially for custom requests in tests.

* **Status Codes**:

  * Use `http.StatusMethodNotAllowed` (405) for disallowed methods like `DELETE` when not supported.
  * Use `http.Error` for client errors (e.g., malformed JSON) and avoid `log.Fatal`, which terminates the server.

* **Testing Improvements**:

  * Prefer `http.NewRequest()` with a custom client in tests for flexibility.
  * Ensure `defer resp.Body.Close()` is always called to prevent resource leaks.
  * Explicitly assert for correct status codes rather than just checking for `!= http.StatusOK`.

### What I learned:

* **`http.Client` Design**: It’s idiomatic in Go to use the pointer form because most methods operate on `*Client`, and the struct holds shared state.
* **Fatal vs Non-Fatal Logging**: `log.Fatal` exits the server and leads to test transport errors; `http.Error` is the proper way to return meaningful HTTP responses.
* **Test Debugging Tactics**: Recognizing the distinction between transport-level failures and status code mismatches is essential for diagnosing test issues.
* **Go HTTP Anatomy**: Learned how to combine server and client concepts—method constants, request creation, and response writing—for robust request handling.

### Next steps:

* Ensure all disallowed methods return 405, not 403, aligning with HTTP semantics.
* Expand the test suite to cover all HTTP methods with clear expectations for each.
* Add JSON tags to the `Message` struct and validate request body parsing for `POST` requests.
* Review and improve all logging and error response patterns for clarity and maintainability.

---

## 📅 [2025-05-22]

### What I worked on:
- Implemented JSON marshaling/unmarshaling with structs for HTTP request handling
- Learned about Go's package organization and struct visibility rules
- Explored dependency injection patterns vs global state for logging
- Debugged issues with Go test caching affecting integration tests
- Set up HTTP handler for processing JSON POST requests

### Problems or blockers:
- Initially struggled with nil pointer dereference when trying to use logger as global variable before initialization
- Encountered confusion with Go test caching making it appear that server wasn't receiving requests

### Decisions made and why:
- **Struct Organization**: Decided to keep related structs in the same package rather than separate files, following Go's "organize by functionality, not by type" principle
- **Global Logger Access**: Chose to stick with global state for logging (controlled global) rather than full dependency injection, balancing convenience vs testability for infrastructure concerns like logging
- **JSON Handling**: Used Go's built-in `encoding/json` package with struct tags for clean serialization/deserialization
- **Testing Approach**: Learned to use `go test -count=1` or `go clean -testcache` to handle integration test caching issues

### What I learned:
- **Go Package System**: Files in the same package can access each other's exported types without imports, but both files must have identical package declarations. You cannot export structs from a test file.
- **io.Reader Interface**: How to convert marshaled JSON byte slices to `io.Reader` using `bytes.NewBuffer()` for HTTP requests
- **Dependency Injection Trade-offs**: Understood the tension between explicit dependencies (testable but verbose) vs global state (convenient but harder to test)
- **Go Test Caching**: How Go aggressively caches test results and why this can be confusing for integration tests that depend on external services

### Next steps:
- Expand on struct learning to include JSON tags
- Begin speccing out the struct for our job que messages that our server will
receive.

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

---

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

