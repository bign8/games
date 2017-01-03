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
