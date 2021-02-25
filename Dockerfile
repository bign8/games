FROM golang:1.16-alpine AS builder

# Build Environment (Cache go assets)
RUN apk add --no-cache upx git
# RUN mkdir -p /go/src/github.com/bign8/games
WORKDIR /app
ENV CGO_ENABLED=0
ENV GO111MODULE=on

# Pull and pre-build dependencies
ADD go.* ./
RUN go mod download
RUN go build -v -installsuffix 'static' -ldflags="-s -w" net/http

# Actually pull in and build the application
ADD . ./
RUN go build -v -installsuffix 'static' -ldflags="-s -w" -o server ./cmd/server

# Compress binary
ARG COMPRESS=true
RUN if [ "$COMPRESS" != "false" ]; then upx --ultra-brute router; fi

# The Actual Container
FROM scratch
COPY --from=builder /app/server /
ENTRYPOINT ["/server"]
EXPOSE 4000
