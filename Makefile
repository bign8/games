serve:
	@gin -a 4000 -t cmd/server -i

install:
	@go get github.com/Masterminds/glide
	@glide install -v

test:
	@go test -race -bench=. -benchmem -v $(shell glide nv)
