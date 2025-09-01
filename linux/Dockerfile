# -----------
# Build Image
# -----------
FROM golang:1.25-alpine3.22 AS builder
ENV CGO_ENABLED=0
#ENV GOOS=linux

WORKDIR /app

# Install source deps
COPY go.mod go.sum ./
RUN go mod download

# Copy source & build
COPY . .

# Build binary:
# -a: force rebuild
# -installsuffix cgo: remove cgo support
# -ldflags="-w -s -extldflags '-static'": enable static linking, omits the DWARF symbol table and debug information
# -trimpath: remove all file system paths from the compiled executable
# -o /helloworld: output binary to /helloworld
RUN go build \
    -a \
    -installsuffix cgo \
    -ldflags="-w -s -extldflags '-static'" \
    -trimpath \
    -o /helloworld .

# Compress with UPX (trade-off: smaller size vs startup time)
RUN apk add --no-cache upx && upx --best --lzma /helloworld

# ------------
# Certificates
# ------------
FROM alpine:3.22 AS certs

RUN apk --no-cache add ca-certificates

# -----------------
# Distributed Image
# -----------------
FROM scratch

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY static /static
COPY --from=builder /helloworld /helloworld
ENTRYPOINT ["/helloworld"]
