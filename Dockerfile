FROM golang:latest
WORKDIR /api-rest
COPY . .
RUN go build -o main .
# CMD ["./main"]
