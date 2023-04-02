FROM golang:alpine

COPY . /

RUN go mod download; \
    go build -mod=vendor -a -o ./bin/svc -ldflags="-s -w"

CMD ["./svc"]
