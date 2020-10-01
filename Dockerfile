FROM golang:1.15
RUN mkdir /go/api-rest
ADD . /go/api-rest
WORKDIR /go/api-rest
RUN go build -o server .
CMD ["/go/api-rest/server"]
