.PHONY: build clean fmt scheck vet
.DEFAULT_GOAL: build
build: test
	go build -o tb -ldflags '-s -w'
clean:
	rm -f tb
fmt:
	go fmt ./...
scheck: vet
	staticcheck ./...
test: scheck
	go test ./...
vet: fmt
	go vet ./...
