# VIPMS1

Simple Go/Fiber VIP visitor management API backed by MySQL. It can create VIP
records, list/search them, and mark a VIP as arrived.

## Requirements

- Go 1.26.5+
- MySQL-compatible database

## Tests

Run local tests without a database:

```bash
make test
```

Run the MySQL integration tests:

```bash
RUN_MYSQL_INTEGRATION=1 make integration-test
```

Run a dependency vulnerability check:

```bash
make vulncheck
```

Run any one-off command through the Makefile:

```bash
make run CMD='go test ./...'
```

## Maintenance Checklist

- [x] Validate VIP name, country, ETA timestamp, and photo URL length before insert.
- [x] Add unit tests for VIP validation without requiring MySQL.
- [x] Update the Go module directive for the current Go toolchain.
- [x] Update the Go runtime directive to Go 1.26.5 for the fixed standard library.
- [x] Refresh Fiber, MySQL, SQLX, and vendored dependencies.
- [x] Gate MySQL integration tests behind `RUN_MYSQL_INTEGRATION=1`.
- [x] Add Makefile targets for tests, integration tests, vulnerability checks, vendoring, and arbitrary commands.
- [x] Run `make test`.
- [x] Run `make vulncheck`; no reachable vulnerabilities were found.
