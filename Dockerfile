FROM golang:1.22.1-bookworm

WORKDIR /app
COPY . .

RUN apt update && apt install -y nodejs
RUN curl -f https://get.pnpm.io/v6.16.js | node - add --global pnpm
RUN make deps
RUN make build

EXPOSE $PORT

ENTRYPOINT ["./bin/server"]
