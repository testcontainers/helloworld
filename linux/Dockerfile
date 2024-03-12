# -----------
# Build Image
# -----------
FROM golang:1.22-alpine3.19 AS builder
ENV CGO_ENABLED=0
#ENV GOOS=linux

WORKDIR /app

# Install source deps
COPY go.mod go.sum ./
RUN go mod download

# Copy source & build
COPY . .
RUN go build -v -ldflags '-extldflags "-static"' -o /helloworld .

# -----------------
# Distributed Image
# -----------------
FROM alpine:3.20
COPY static /static
COPY --from=builder /helloworld /helloworld
ENTRYPOINT ["/helloworld"]
