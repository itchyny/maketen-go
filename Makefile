BIN := maketen
VERSION = $$(make show-version)
CURRENT_REVISION := $(shell git rev-parse --short HEAD)
BUILD_LDFLAGS := "-X github.com/itchyny/maketen-go/cli.revision=$(CURRENT_REVISION)"
export GO111MODULE=on

.PHONY: all
all: clean build

.PHONY: build
build: deps
	go build -ldflags=$(BUILD_LDFLAGS) -o build/$(BIN) ./cmd/$(BIN)

.PHONY: install
install: deps
	go install -ldflags=$(BUILD_LDFLAGS) ./...

.PHONY: deps
deps:
	go get -d -v ./...

.PHONY: show-version
show-version:
	@GO111MODULE=off go get github.com/motemen/gobump/cmd/gobump
	@gobump show -r cli

.PHONY: cross
cross: crossdeps
	goxz -n $(BIN) -pv=v$(VERSION) -build-ldflags=$(BUILD_LDFLAGS) ./cmd/$(BIN)

.PHONY: crossdeps
crossdeps: deps
	GO111MODULE=off go get github.com/Songmu/goxz/cmd/goxz

.PHONY: test
test: build
	go test -v ./...

.PHONY: lint
lint: lintdeps
	go vet ./...
	golint -set_exit_status ./...

.PHONY: lintdeps
lintdeps:
	GO111MODULE=off go get golang.org/x/lint/golint

.PHONY: clean
clean:
	rm -rf build goxz
	go clean

.PHONY: bump
bump:
	@git status --porcelain | grep "^" && echo "git workspace is dirty" >/dev/stderr && exit 1 || :
	gobump set $(shell sh -c 'read -p "input next version (current: $(VERSION)): " v && echo $$v') -w cli
	git commit -am "bump up version to $(VERSION)"
	git tag "v$(VERSION)"
	git push
	git push --tags

.PHONY: crossdocker
crossdocker:
	docker run --rm -v `pwd`:"/$${PWD##*/}" -w "/$${PWD##*/}" golang make cross

.PHONY: upload
upload:
	GO111MODULE=off go get github.com/tcnksm/ghr
	ghr "v$(VERSION)" goxz

.PHONY: release
release: test lint clean bump crossdocker upload
