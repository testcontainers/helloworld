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
RUN go build -v -ldflags '-extldflags "-static"' -o /helloworld .

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
