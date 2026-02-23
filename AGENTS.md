# Go Deepcopy

## Purpose

Go Deepcopy provides utilities for deep copying Go objects. It enables safe duplication of complex data structures used across microservices without shared memory references.

## Type

Shared library

## Consumed By

Services requiring object duplication:
- backtest-manager-service
- strategy-service, rules-service
- report-service
- data-manager-service

## Exposed Interface

**Key Functions:**

- `DeepCopy(src, dst interface{})`: Copy src to dst recursively
- `Clone(v interface{}) interface{}`: Create clone of value
- `DeepCopySlice(src []interface{}) []interface{}`: Copy slice with deep elements
- `DeepCopyMap(src map[string]interface{}) map[string]interface{}`: Copy map deeply

## Usage Example

```go
import "github.com/arconomy/go-deepcopy"

original := &Strategy{Name: "Test", Rules: []*Rule{...}}
copy := deepcopy.Clone(original).(*Strategy)
```

## MCP Rules

**When This Library Changes, Services That Must Be Updated**:
- Any service using Clone or DeepCopy functions
- Update via `go get -u github.com/arconomy/go-deepcopy`

**Testing Changes**:
1. Run `go test ./...` in go-deepcopy
2. Verify cloned objects are independent in consuming services

---

**Last Updated**: 2026-02-22



After ANY change to this service that affects:
- API surface (new/modified endpoints or gRPC methods)
- NATS subjects (new publish/subscribe calls)
- Data models (schema/migration changes)
- Config vars (new env vars)
- Inter-service dependencies (new gRPC clients or NATS consumers)

You MUST update this service's AGENTS.md to reflect the change before committing.
Update the root ./AGENTS.md if the change affects the architecture graph,
NATS subject registry, or gRPC relationship table.

---

**Last Updated**: 2026-02-23
