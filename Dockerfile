FROM golang:1.16
WORKDIR /tmp/app
COPY . .
RUN go build -o main .
CMD ["/tmp/app/main"]