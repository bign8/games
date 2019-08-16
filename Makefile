SHELL:=/bin/bash -o pipefail

make: bench
.PHONY : make

test:
	go test -race ./...
.PHONY : test

bench:
	go install github.com/bign8/gobench
	go test -race -bench=. -benchmem -v ./... | tee test.out
	gobench -in test.out
.PHONY : bench

serve:
	go run cmd/server/main.go
.PHONY : serve

build:
	GO111MODULE=on go build ./cmd/server
.PHONY : build

docker:
	docker build -t bign8/games --build-arg COMPRESS=false .
	docker run --rm -it -p 4000:4000 bign8/games
.PHONY : docker

watch:
	go run vendor/github.com/codegangsta/gin/main.go -a 4000 -d cmd/server -i --all -- -tout 1ms
.PHONY : watch
