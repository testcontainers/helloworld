ARG BASE_IMAGE

# -----------
# Build Image
# -----------
FROM golang:1.25-nanoserver as builder
ENV CGO_ENABLED=0

WORKDIR /app

# Install source deps
COPY go.mod go.sum ./
RUN go mod download

# Copy source & build
COPY . .
RUN go build -v -ldflags "-s" -o /bin/helloworld

# -----------------
# Distributed Image
# -----------------
FROM ${BASE_IMAGE}
COPY static /static
COPY --from=builder /bin/helloworld /bin/helloworld
ENTRYPOINT ["/bin/helloworld"]
