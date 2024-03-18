FROM golang:1.22.1-bookworm

WORKDIR /app

COPY go.* .

RUN go mod tidy
RUN go mod verify
RUN go mod download

COPY main.go .

RUN go build -o app .

EXPOSE $PORT

ENTRYPOINT ["./app"]
