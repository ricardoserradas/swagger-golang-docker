# Containerizing this Go REST API

If you look at the Dockerfile, you'll see that I first use the `golang` image to build the code and then I store only the binaries on a `scratch` image. This is because if we create this image with the `golang` image, we're going to get an ~800MB final image.

With the `scratch` image to store only the binaries, we're getting an ~5MB image.

## Generate static Go binaries

Be aware that, to make the binary executable on the scratch image, which has no runtimes installed, I needed to build the source with the following modifiers:

`CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w -extldflags "-static"' `

What means:

* `CGO_ENABLED=0`: enables it to be built to run in multiple platforms;
* `GOOS=linux`: enables it to run on linux;
* `-a`: force rebuilding of packages that are already up-to-date.
* `-w`: Omit the DWARF symbol table;
* `-extldflags "-static"`: to create a static binary.

This will produce a binary which we can run standalone in a clean distro.

## Building the container

`docker build -t goswagger-training .`

## Running the container

`docker run --publish 6060:8080 --name goswagger-training --rm goswagger-training`

## References

### Golang Docker image
A Debian Linux container with Go development environment, for building purposes.
https://hub.docker.com/_/golang/

### Scratch Docker image
An explicit empty image, especially to create very tiny images.
https://hub.docker.com/_/scratch/

### Deploying Go Servers with Docker
https://blog.golang.org/docker

### Building Docker images for Static Go binaries
https://medium.com/@kelseyhightower/optimizing-docker-images-for-static-binaries-b5696e26eb07

### Totally static Go builds
http://blog.wrouesnel.com/articles/Totally%20static%20Go%20builds/