FROM golang:latest
RUN mkdir /api-rest
ADD . /api-rest
WORKDIR /api-rest
RUN go build -o server .
CMD ["./server", "-rt"]
