POSTGRESQL_URL=host='' port=5432 user='' password='' dbname=''
LOG_LEVEL=debug

.PHONY: vendor
vendor:
	GO111MODULE=${GO111MODULE} go get ./... && \
	GO111MODULE=${GO111MODULE} go mod tidy && \
	GO111MODULE=${GO111MODULE} go mod vendor

.PHONY: migration
migration:
	POSTGRESQL_URL="${POSTGRESQL_URL}" go run cmd/migration/main.go

.PHONY: u-test
u-test:
	go test -v -cover -count=1 -mod vendor ./internal/...; \

.PHONY: i-test
i-test:
	POSTGRESQL_URL="${POSTGRESQL_URL}" go test -v -cover -tags=integration -count=1 -mod vendor ./...; \

.PHONY: e2e-test
e2e-test:
	go test -v -cover -tags=e2e -count=1 -mod vendor ./...; \

.PHONY: lint
lint:
	golangci-lint run
