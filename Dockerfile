FROM golang:1.14-alpine

ADD . /payment-system

WORKDIR /payment-system

RUN make build OUTPUT=/go/bin/payment-system

EXPOSE 8080

ENTRYPOINT ["/go/bin/payment-system", "POSTGRESQL_URL={POSTGRESQL_URL}"]
