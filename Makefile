SHELL:=/bin/bash -o pipefail

make: vendor bench
.PHONY : make

test:
	go test -race $(shell glide nv)
.PHONY : test

bench:
	go test -race -bench=. -benchmem -v $(shell glide nv) | tee test.out
	gobench -in test.out
.PHONY : bench

vendor: glide.lock
	go get github.com/Masterminds/glide
	glide install -v
	go install ./vendor/github.com/bign8/gobench

serve:
	go run cmd/server/main.go
.PHONY : serve

build:
	go build ./cmd/server
.PHONY : build

# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o games

docker:
	GOOS=linux GOARCH=amd64 go build -i -v -ldflags "-s -w" -installsuffix cgo ./cmd/server
	docker build -t bign8/games .
	docker run --rm -it -p 4000:4000 bign8/games
.PHONY : docker
