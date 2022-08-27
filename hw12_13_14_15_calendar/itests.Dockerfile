# Собираем в гошке
FROM golang:1.17.6 as build

RUN mkdir -p /opt/integration_tests
WORKDIR /opt/integration_tests

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

ENTRYPOINT ["go", "test", "./itests/...", "-v"]


