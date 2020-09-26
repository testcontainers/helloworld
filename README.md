# Testcontainers Helloworld Docker Image

This is a Docker image for use by Testcontainers' own self-test suites. It is not intended for use outside of the Testcontainers project.

It features a small HTTP server with the following characteristics:

* It serves content on two ports (8080 and 8081) to enable testing that multiple Docker container ports can be exposed.
* It serves an HTML root page, with a few basic elements, to enable verification that browser-based test tools can access the container.
* It servers a non-HTML endpoint at `/ping`.
* It implements a configurable delay at startup before each port's server is started, to enable testing of startup wait strategies (TCP or HTTP-based). Setting the environment variable `DELAY_START_MSEC` to a non-zero number will:
    * wait for the defined duration
    * start the port 8080 server
    * wait again for the same duration
    * start the port 8081 server
* It emits a basic log message after starting which can be used to test log-based wait strategies.

## Example usage

```
$ docker run -p 8080:8080 -p 8081:8081 -e DELAY_START_MSEC=2000 helloworld

2020/09/26 08:50:55 DELAY_START_MSEC: 2000
2020/09/26 08:50:55 Sleeping for 2000 ms
2020/09/26 08:50:57 Starting server on port 8080
2020/09/26 08:50:57 Sleeping for 2000 ms
2020/09/26 08:50:59 Starting server on port 8081
2020/09/26 08:50:59 Ready, listening on 8080 and 8081
```

## License

See [LICENSE](./LICENSE).

## Copyright

Copyright (c) 2020 Richard North and other authors.