.PHONY: build clean fmt scheck vet
build: vet
	go build -o tb -ldflags '-s -w'
clean:
	rm -f tb
fmt:
	go fmt ./...
scheck:
	staticcheck ./...
vet: fmt
	go vet ./...
