POSTGRESQL_URL=

.PHONY: vendor
vendor:
	GO111MODULE=${GO111MODULE} go get ./... && \
	GO111MODULE=${GO111MODULE} go mod tidy && \
	GO111MODULE=${GO111MODULE} go mod vendor

.PHONY: migration
migration:
	POSTGRESQL_URL="${POSTGRESQL_URL}" go run cmd/migration/main.go
