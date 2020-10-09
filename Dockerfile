FROM golang:1.14-alpine

ADD . /payment-system

WORKDIR /payment-system

RUN make build OUTPUT=/go/bin/payment-system

ENV POSTGRESQL_URL="host='10.233.33.234' port=5432 user='payment_system' password='payment_system' dbname='payment_system'"

EXPOSE 8080

ENTRYPOINT ["/go/bin/payment-system", "POSTGRESQL_URL={POSTGRESQL_URL}"]
