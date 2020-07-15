POSTGRESQL_URL=host='XX.XXX.XX.XXX' port=5432 user='' password='' dbname=''

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
	POSTGRESQL_URL="${POSTGRESQL_URL}" go test -v -cover -count=1 -mod vendor ./internal/persistence/...; \
