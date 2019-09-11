.PHONY: all clean image test

all:
	go build -o dependency-action cmd/dependency-action/main.go

clean:
	rm -f dependency-action

image:
	docker build .

test:
	ginkgo acceptance

