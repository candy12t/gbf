BIN := bin/gbf
BUILD_LDFLAGS := "-s -w"

.PHONY: all
all: clean build

.PHONY: build
build:
	go build -trimpath -ldflags=$(BUILD_LDFLAGS) -o $(BIN) .

.PHONY: install
install:
	go install -ldflags=$(BUILD_LDFLAGS) .

.PHONY: clean
clean:
	rm -rf $(BIN)
	go clean

.PHONY: test
test:
	go test -race ./...
