FROM golang:1.20

WORKDIR /usr/src/app

COPY . .

RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go mod download && go mod verify

RUN go build -v -o ./ ./...

CMD ["./question-bot"]