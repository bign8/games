SHELL:=/bin/bash -o pipefail
.PHONY:= make test bench install serve build docker

make: install bench

test:
	go test -race $(shell glide nv)

bench:
	go test -race -bench=. -benchmem -v $(shell glide nv) | tee test.out
	gobench -in test.out

install:
	go get github.com/Masterminds/glide
	glide install -v
	go install -v ./vendor/github.com/bign8/gobench
	rm -rf ${GOPATH}/src/github.com/gopherjs
	go install -v ./vendor/github.com/gopherjs/gopherjs
	mv vendor/github.com/gopherjs ${GOPATH}/src/github.com/gopherjs

serve:
	gin -a 4000 -t cmd/server -i

build:
	go build ./cmd/server

# vendor: glide.yaml glide.lock ## Fetch go vendored dependencies
# 	rm -rf ${GOPATH}/src/github.com/gopherjs
# 	mv vendor/github.com/gopherjs ${GOPATH}/src/github.com/gopherjs
#
# glide.yaml:
# 	@# TODO: get this to skip glide.lock changes
# 	glide update --strip-vendor
#
# glide.lock: glide.yaml
# 	glide install --strip-vendor

# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o games

docker:
	GOOS=linux GOARCH=amd64 go build -i -v -ldflags "-s -w" -installsuffix cgo ./cmd/server
	docker build -t bign8/games .
	docker run --rm -it -p 4000:4000 bign8/games
