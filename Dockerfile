FROM golang:1.14.2-alpine3.11 as builder

ENV GO111MODULE=on

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
  go build -v -mod=vendor -o dns

# Copy compiled binary to clear Alpine Linux image
FROM alpine:3.11.5
WORKDIR /
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/mxssl/dns/dns /usr/local/bin/dns
RUN chmod +x /usr/local/bin/dns