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
# Build binary:
# -a: force rebuild
# -installsuffix cgo: remove cgo support
# -ldflags="-w -s -extldflags '-static'": enable static linking, omits the DWARF symbol table
# -trimpath: remove all file system paths from the compiled executable
# -o /bin/helloworld: output binary to /bin/helloworld
RUN go build \
    -a \
    -installsuffix cgo \
    -ldflags="-w -s -extldflags '-static'" \
    -trimpath \
    -o /bin/helloworld .

# Do not optimise with UPX on Windows, as nanoserver does not have PowerShell to install it

# -----------------
# Distributed Image
# -----------------
FROM ${BASE_IMAGE}
COPY static /static
COPY --from=builder /bin/helloworld /bin/helloworld
ENTRYPOINT ["/bin/helloworld"]
