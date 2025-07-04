FROM golang:1.21.5-alpine3.17 as builder

ENV GO111MODULE=on
ARG VERSION=dev

WORKDIR /go/src/github.com/mxssl/dns
COPY . .

# Install external dependcies
RUN apk add --no-cache \
  ca-certificates \
  curl \
  git

# Compile binary

RUN CGO_ENABLED=0 \
  GOOS=`go env GOHOSTOS` \
  GOARCH=`go env GOHOSTARCH` \
  go build -v -o dns -ldflags "-X github.com/mxssl/dns/cmd.version=$VERSION"

# Copy compiled binary to clear Alpine Linux image
FROM alpine:3.22.0
WORKDIR /
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/mxssl/dns/dns /usr/local/bin/dns
RUN chmod +x /usr/local/bin/dns
