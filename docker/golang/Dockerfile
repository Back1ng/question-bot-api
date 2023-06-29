FROM golang:1.20

WORKDIR /usr/src/app

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . .

RUN go mod download && go mod verify

RUN go build -v -o ./ ./...

ENTRYPOINT [ "tail", "-f", "/dev/null" ]