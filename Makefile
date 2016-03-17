PROGRAM=jagger
BUILD_DIR=_build

all: install

install: deps
	go install

deps:
	go get

build: clean-build linux-arm64 mac-amd64

linux-arm64: deps
	GOOS=linux GOARCH=arm64 go build -o $(BUILD_DIR)/$@/$(PROGRAM)

mac-amd64: deps
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$@/$(PROGRAM)

clean: clean-build
	rm -f $(PROGRAM)

clean-build:
	rm -rf $(BUILD_DIR)

.PHONY: install build clean
