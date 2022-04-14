BIN := "./bin/calendar"
DOCKER_IMG="calendar:develop"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/calendar

run: build
	$(BIN) -config ./configs/config.toml

build-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_IMG) \
		-f build/Dockerfile .

run-img: build-img
	docker run $(DOCKER_IMG)

version: build
	$(BIN) version

test:
	go test -race ./internal/...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.41.1

lint: install-lint-deps
	golangci-lint run ./...

migrate:
	goose postgres "host=0.0.0.0 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable" up

.vendor-proto:
	#mkdir -p vendor.protogen
	#mkdir -p vendor.protogen/api/calendar-event
	cp api/proto/EventService.proto vendor.protogen/api/calendar-event

#	@if [ ! -d vendor.protogen/google ]; then \
#		git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
#		mkdir -p  vendor.protogen/google/ &&\
#		mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
#		rm -rf vendor.protogen/googleapis ;\
#	fi
#	@if [ ! -d vendor.protogen/google/protobuf ]; then \
#	        git clone https://github.com/protocolbuffers/protobuf vendor.protogen/protobuf &&\
#	        mkdir -p  vendor.protogen/google/protobuf &&\
#	        mv vendor.protogen/protobuf/src/google/protobuf/*.proto vendor.protogen/google/protobuf &&\
#		rm -rf vendor.protogen/protobuf ;\
#	fi
#	@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
#		mkdir -p vendor.protogen/validate &&\
#		git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
#		mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
#		rm -rf vendor.protogen/protoc-gen-validate ;\
#	fi

generate-proto:
	PATH="${PATH}:${HOME}/go/bin" protoc -I vendor.protogen \
    	--go_out=pkg/calendar-event --go_opt=paths=import \
        --go-grpc_out=pkg/calendar-event --go-grpc_opt=paths=import \
        --grpc-gateway_out=pkg/calendar-event \
        --grpc-gateway_opt=logtostderr=true \
        --grpc-gateway_opt=paths=import \
        api/calendar-event/EventService.proto

evans-connect:
	evans --path vendor.protogen --proto api/calendar-event/EventService.proto repl -p 8081

.PHONY: build run build-img run-img version test lint
