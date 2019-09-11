.PHONY: all clean ginkgo image test

all:
	go build -o dependency-action cmd/dependency-action/main.go

clean:
	rm -f dependency-action

image:
	docker build .

test: ginkgo
	$(GINKGO) acceptance

# download ginkgo if necessary
ginkgo:
ifeq (, $(shell which ginkgo))
	go get github.com/onsi/ginkgo/ginkgo
GINKGO=$(GOBIN)/ginkgo
else
GINKGO=$(shell which ginkgo)
endif
