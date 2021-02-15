FROM golang:1.15.8

WORKDIR /go/src/app/userservice

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY ./src ./src

CMD ["go", "run", "./src/server.go"]