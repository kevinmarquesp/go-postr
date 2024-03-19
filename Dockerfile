FROM golang:1.22.1-bookworm

WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o app .

EXPOSE $PORT

ENTRYPOINT ["./app"]
