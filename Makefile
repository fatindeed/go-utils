all: test

GIT_TAG := $(shell git describe --tags --abbrev=0)

test:
	go test -cover -v ./...

publish:
	git push
	git push --tags
	curl https://proxy.golang.org/github.com/fatindeed/go-utils/@v/$(GIT_TAG).info