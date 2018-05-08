FROM golang

ADD . /go/src/goswagger-training
WORKDIR /go/src/goswagger-training

RUN go get -d -v
RUN go install
RUN GOOS=linux GOARCH=amd64
RUN go build

ENTRYPOINT [ "goswagger-training" ]

EXPOSE 8080

# FROM golang:1.8 as goimage
# ENV SRC=/go/src

# RUN mkdir -p /go/src

# WORKDIR /go/src/app
# COPY . .

# RUN go get -d -v
# RUN go install
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
# go build -o bin/goswagger

# FROM alpine:3.6 as baseimagealp
# RUN apk add --no-cache bash
# ENV WORK_DIR=/docker/bin
# WORKDIR $WORK_DIR
# COPY --from=goimage /go/src/app/bin/ ./

# ENTRYPOINT /docker/bin/goswagger
# EXPOSE 8080