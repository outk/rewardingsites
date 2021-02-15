FROM golang:1.15.8

WORKDIR /go/src/app/userservice

COPY ./src/userservice/go.mod ./src/userservice/go.sum ./

RUN go mod download

COPY ./src/userservice/src ./src

CMD ["go", "run", "./src/server.go"]