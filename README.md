# Golang Advanced Concepts

A hands-on reference for Go internals and production patterns — each topic is a self-contained, runnable example covering the concept, common pitfalls, and idiomatic usage.

---

## Topics Covered

### Concurrency
| Topic | What it covers |
|-------|----------------|
| [goroutines](./goroutines) | Spawning goroutines, lifecycle management, leak prevention |
| [channels](./channels) | Buffered vs unbuffered, directional channels, closing semantics |
| [select](./select) | Multi-channel select, default case, timeout patterns |
| [worker_pool](./worker_pool) | Fan-out/fan-in, bounded concurrency with a fixed goroutine pool |
| [mutex](./mutex) | `sync.Mutex` vs `sync.RWMutex`, protecting shared state |
| [atomic](./atomic) | Lock-free counters with `sync/atomic`, when to prefer over mutex |
| [racecondition](./racecondition) | Detecting data races with `-race`, common causes and fixes |
| [workflows](./workflows) | Pipeline patterns, staged processing, cancellation propagation |

### Error Handling
| Topic | What it covers |
|-------|----------------|
| [errors](./errors) | `errors.Is`, `errors.As`, wrapping with `%w`, sentinel errors |
| [Error handling](./Error%20handling) | Idiomatic patterns, error types, avoiding panic for control flow |
| [defer_panic_recover](./defer_panic_recover) | Defer execution order, panic/recover for boundary protection |

### Interfaces & Types
| Topic | What it covers |
|-------|----------------|
| [Interface](./Interface) | Implicit implementation, polymorphism, interface composition |
| [Stringer_Interface](./Stringer_Interface) | Implementing `fmt.Stringer` for custom string output |
| [Writer_Interface](./Writer_Interface) | `io.Writer` pattern, composing writers |
| [type_assertion](./type_assertion) | Type switches, safe assertions, interface to concrete type |
| [embedding](./embedding) | Struct and interface embedding, promoting methods |
| [generics](./generics) | Type parameters, constraints, generic data structures |

### Context & Lifecycle
| Topic | What it covers |
|-------|----------------|
| [context](./context) | Deadlines, cancellation, passing values — `WithTimeout`, `WithCancel` |
| [closures](./closures) | Closure over variables, gotchas in goroutines, functional patterns |

### Serialization
| Topic | What it covers |
|-------|----------------|
| [marshal](./marshal) | JSON encoding, struct tags, custom `MarshalJSON` |
| [unMarshal](./unMarshal) | JSON decoding, handling unknown fields, custom `UnmarshalJSON` |

### Testing
| Topic | What it covers |
|-------|----------------|
| [testing](./testing) | Table-driven tests, subtests, benchmarks, `testify` usage |

---

## Running Examples

Each topic directory contains a `main.go` or `*_test.go`. Run any example with:

```bash
go run ./goroutines
go run ./worker_pool
go test ./testing/...
```

To check for race conditions:

```bash
go run -race ./racecondition
```

---

## Prerequisites

- Go 1.22+
- No external dependencies — stdlib only (except where noted)

---

## Who this is for

Engineers who know basic Go and want to understand the internals well enough to make confident decisions in production — when to use a mutex vs atomic, how to avoid goroutine leaks, how generics compose with interfaces.
