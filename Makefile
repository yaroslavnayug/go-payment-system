OUTPUT?=bin/payment-system
POSTGRESQL_URL?=host='0.0.0.0' port=5432 user='root' password='root' dbname='payment_system'
LOG_LEVEL=debug

.PHONY: vendor
vendor:
	GO111MODULE=${GO111MODULE} go get ./... && \
	GO111MODULE=${GO111MODULE} go mod tidy && \
	GO111MODULE=${GO111MODULE} go mod vendor

.PHONY: u-test
u-test:
	GO111MODULE=on go test -v -mod=vendor -cover ./... -coverprofile cover.out;

.PHONY: i-test
i-test:
	GO111MODULE=on POSTGRESQL_URL="${POSTGRESQL_URL}" go test ./internal/postgres -tags=integration -v -mod=vendor -cover -coverprofile cover.out;

.PHONY: e2e-test
e2e-test:
	GO111MODULE=on go test -tags=e2e -v -mod=vendor -count=1  ./test/e2e; \

.PHONY: coverage
coverage:
	@echo "+ $@"
	GO111MODULE=on go tool cover -func cover.out | grep total | awk '{print $3}'

.PHONY: lint
lint:
	golangci-lint run

.PHONE: swagger
swagger:
	swagger generate spec -o ./docs/swagger.json

.PHONE: build
build:
	GO111MODULE=${GO111MODULE} go build \
    		-mod vendor \
    		-o ${OUTPUT} cmd/server/main.go
