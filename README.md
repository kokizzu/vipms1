# VIPMS1

Simple Go/Fiber VIP visitor management API backed by MySQL. It can create VIP
records, list/search them, and mark a VIP as arrived.

## Requirements

- Go 1.26+
- MySQL-compatible database

## Tests

Run the validation unit tests without a database:

```bash
go test ./app/model -run TestVipValidation
```

The older integration tests still expect a local MySQL instance configured by
the app config.

## Maintenance Checklist

- [x] Validate VIP name, country, ETA timestamp, and photo URL length before insert.
- [x] Add unit tests for VIP validation without requiring MySQL.
- [x] Update the Go module directive for the current Go toolchain.
