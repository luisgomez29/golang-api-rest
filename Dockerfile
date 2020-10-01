FROM golang:latest
RUN mkdir /go/api-rest
ADD . /go/api-rest
WORKDIR /go/api-rest
RUN go build -o server .
CMD ["/go/api-rest/server"]
