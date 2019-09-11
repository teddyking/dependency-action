.PHONY: all clean ginkgo image test

all:
	go build -o dependency-action cmd/dependency-action/main.go

clean:
	rm -f dependency-action

image:
	docker build .

test:
	go get github.com/onsi/ginkgo/ginkgo
	ginkgo acceptance
