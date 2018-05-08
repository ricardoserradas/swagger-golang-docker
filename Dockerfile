FROM golang

ADD . /go/src/goswagger-training
WORKDIR /go/src/goswagger-training

RUN go get -d -v
RUN go install
RUN GOOS=linux GOARCH=amd64
RUN go build

ENTRYPOINT [ "goswagger-training" ]

EXPOSE 8080