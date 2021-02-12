FROM golang:latest

COPY ./src /src

CMD ["go", "run", "/src/main.go"]