make: install bench

test:
	go test -race $(shell glide nv)

bench:
	go test -race -bench=. -benchmem -v $(shell glide nv) | tee test.out
	gobench -in test.out

install:
	go get github.com/Masterminds/glide
	glide install -v
	go install ./vendor/github.com/bign8/gobench

serve:
	gin -a 4000 -t cmd/server -i

build:
	go build ./cmd/server

docker:
	GOOS=linux GOARCH=amd64 go build -i -v -ldflags "-s -w" -installsuffix cgo ./cmd/server
	docker build -t bign8/games .
	docker run --rm -it -p 4000:4000 bign8/games
