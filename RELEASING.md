# Releasing Helloworld

This is a Docker image for use by Testcontainers' own self-test suites. It is not intended for use outside the Testcontainers project.
So releases are published as Docker image to https://hub.docker.com/r/testcontainers/hellworld.

The `.github/workflows/publish-docker-image.yml` workflow is our primary release workflow,
while preview images are published via `.github/workflows/build-docker-image.yml`.

## Multi-Architecture support

The workflows publishing the `testcontainers/helloworld` Docker image ensure that multiple platforms are supported.
Supported platforms are listed in the "supported-architectures.txt" file. The file is used as reference for
the checks after publishing the multi-arch image.
