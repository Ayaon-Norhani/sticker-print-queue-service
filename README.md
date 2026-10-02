# 🏷️ Sticker & Label Print Queue Service

A high-performance, autonomous Go microservice designed to validate, process, and queue custom print orders (stickers, die-cut vinyl, and labels) for high-volume manufacturing pipelines.

Built with idiomatic Go patterns, context-aware PostgreSQL queries, and AI-accelerated unit testing workflows.

---

## ✨ Features

- **Guard-Clause Validation:** Utilizes early returns and custom domain-specific sentinel errors (`ErrInvalidSKU`, `ErrQuantityOutOfRange`, `ErrInvalidMaterial`) to enforce business rules cleanly.
- **Context-Aware Database Layer:** Implements `context.Context` timeout handling for PostgreSQL executions to prevent thread blocking and ensure database safety.
- **Table-Driven Unit Testing:** High test coverage utilizing Go's native `testing` framework to validate happy paths, edge cases, and invalid payloads.
- **AI-Accelerated Engineering:** Leveraged AI developer tools to rapidly generate table-driven test cases and optimize query execution logic.

---

## 🛠️ Tech Stack

- **Language:** Go (Golang 1.22+)
- **Database:** PostgreSQL
- **Testing:** Native Go `testing` package
- **Architecture:** Clean Architecture / Command Pattern Validation

---

## 🚀 Getting Started

### Prerequisites

- Go 1.22 or higher installed
- PostgreSQL (optional for local running, mocked in unit tests)

### Running Unit Tests

Execute all table-driven unit tests with verbose output:

```bash
go test -v ./...
