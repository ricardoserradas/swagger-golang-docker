# Creating an intermediate image to build the code to linux
FROM golang as buildimage

ADD . /go/src/goswagger-training
WORKDIR /go/src/goswagger-training

RUN go get -d -v
RUN go install

# Building it as static as possible so it can run on the scratch image
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w -extldflags "-static"' .

# Put binaries into a thinner image
FROM scratch as baseimage
WORKDIR /docker/bin
COPY --from=buildimage /go/src/goswagger-training/goswagger-training ./
ENV PORT 8080
EXPOSE 8080

ENTRYPOINT [ "/docker/bin/goswagger-training" ]