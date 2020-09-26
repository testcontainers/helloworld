FROM golang:alpine AS builder
WORKDIR /go/src/github.com/testcontainers/helloworld
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /helloworld .

FROM scratch
COPY static /static
COPY --from=builder /helloworld /helloworld
CMD ["/helloworld"]
