
GOPATH:=$(shell go env GOPATH)

.PHONY: build
build:
	go build -o ./bin/micro main.go plugin.go

.PHONY: build_linux
build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-w' -o ./bin/linux_amd64/micro main.go plugin.go

.PHONY: docker
docker: build_linux
	docker build . -t $(tag)

.PHONY: run
run:
	test -n "$(profile)"
	test -n "$(env)"
	./bin/micro --profile=$(profile) --env=$(env) gateway --address=$(addr)
