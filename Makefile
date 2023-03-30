.PHONY: build clean fmt scheck vet
build: scheck
	go build -o tb -ldflags '-s -w'
clean:
	rm -f tb
fmt:
	go fmt ./...
scheck: vet
	staticcheck ./...
vet: fmt
	go vet ./...
