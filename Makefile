GO ?= go
GOVULNCHECK ?= govulncheck
CMD ?=

.PHONY: test integration-test vulncheck verify-dependency-security vendor run

test:
	$(GO) test ./...

integration-test:
	RUN_MYSQL_INTEGRATION=1 $(GO) test -v ./...

vulncheck:
	$(GOVULNCHECK) ./...

verify-dependency-security: vulncheck

vendor:
	$(GO) mod vendor

run:
	@test -n "$(CMD)" || (echo "usage: make run CMD='go test ./...'" >&2; exit 2)
	$(CMD)
