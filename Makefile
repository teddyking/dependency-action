.PHONY: all clean test

all:
	go build -o dependency-action cmd/dependency-action/main.go

clean:
	rm -f dependency-action

test:
	ginkgo acceptance

