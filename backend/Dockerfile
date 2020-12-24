FROM golang:1.15

WORKDIR /go/src/app

COPY go.mod .

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GO_ENV=production

COPY ./ ./

RUN go build -o main.exe

EXPOSE 9090

CMD ["./main.exe"]
