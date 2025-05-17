# Go Job Queue (HTTP-based Worker System)

This project is a minimal HTTP job queue and worker pool written entirely in Go using only the standard library. It is part of a personal learning initiative to deepen my backend engineering skills and broaden my familiarity with Go as a systems programming language.

The server accepts jobs via HTTP, queues them for processing, and allows clients to check the status or result of a given job. All functionality, including concurrency, HTTP handling, and internal state management, is implemented from scratch to build a deeper understanding of back-end systems programming.

---

## 🚀 Project Goal

To build a fully functioning, concurrent, HTTP-based job queue that:
- Accepts new jobs via POST requests
- Assigns jobs to worker goroutines for processing
- Tracks and stores job statuses and results
- Allows clients to query the status and result of submitted jobs
- Uses no third-party dependencies (standard library only)

---

## 🎯 Core Learning Goals

This project is designed to solidify foundational backend engineering skills and give me fluency in Go’s systems-level constructs.

### 🧵 Concurrency
- Create and manage a pool of workers using goroutines and channels
- Ensure safe access to shared state using synchronization primitives (`sync.Mutex`, `sync.Map`)
- Handle job dispatching and coordination using channels

### 🌐 Networking
- Build a basic HTTP API using `net/http`
- Understand and apply HTTP methods, routing, status codes, and content negotiation
- Implement JSON request and response handling via `encoding/json`

### 🧠 Systems Design
- Implement an in-memory job queue and result store
- Generate unique identifiers for job tracking
- Track job lifecycle: `queued`, `in-progress`, `done`, `failed`
- Manage graceful shutdown and resource cleanup with `context` and `os/signal`

### 🔍 Observability (optional extension)
- Introduce logging to trace job execution
- Design optional metrics around throughput, latency, and job state counts

---

## 🛤 Project Phases

This project is broken into clean, progressive milestones:

1. **Scaffold basic HTTP server**
2. **Define job struct and status enums**
3. **Create in-memory job store**
4. **Add `/job` POST handler to accept new jobs**
5. **Add `/status/{id}` GET handler to report job status**
6. **Implement job queue and worker pool**
7. **Tie worker pool to job store and update job lifecycle**
8. **Add graceful shutdown and signal handling**
9. *(Optional)* Add `/result/{id}` endpoint
10. *(Optional)* Add retry, timeout, or persistent storage layers

---

## 🔒 Constraints

- No third-party libraries or frameworks
- All concurrency handled manually via goroutines and channels
- Entirely in-memory — persistence is out of scope for the initial version
- Keep the scope limited to emphasize backend system mechanics over product features

---

## 📎 Why This Project?

This project is intentionally focused on **backend fundamentals** rather than full-stack integration. While I have strong front-end experience, this serves as a learning exercise to:
- Gain confidence in Go as a backend tool
- Practice concurrent programming in a real-world context
- Build system-level intuition around job scheduling and processing
- Prepare for more advanced backend systems (distributed queues, schedulers, persistent workers, etc.)

---

## 🗺️ What's Next?

Each phase builds logically on the last. I will update this README or create a `DEVLOG.md` to track progress, open questions, and refactoring insights.
