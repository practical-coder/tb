.PHONY: build clean fmt vet
build: vet
	go build -o tb -ldflags '-s -w'
clean:
	rm -f tb
fmt:
	go fmt ./...
vet: fmt
	go vet ./...
