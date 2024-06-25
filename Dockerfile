FROM golang:1.18-alpine3.16 AS builder
WORKDIR /go/src/github.com/testcontainers/helloworld
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /helloworld .

FROM alpine:3.20
COPY static /static
COPY --from=builder /helloworld /helloworld
ENTRYPOINT ["/helloworld"]
