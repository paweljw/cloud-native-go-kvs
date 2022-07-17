test:
	cd pkg/kvs && go test
.PHONY: test

run:
	go run cmd/main.go
.PHONY: run

build:
	go build -o bin/kvs cmd/main.go

clean:
	rm -f bin/kvs
