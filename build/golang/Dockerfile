FROM golang:1.20

RUN apt-get update && \
    apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent

RUN curl -sSL https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
RUN echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ bionic main" > /etc/apt/sources.list.d/migrate.list
RUN apt-get update && \
    apt-get install -y migrate
    
WORKDIR /usr/src/app

COPY . .

RUN go mod download && go mod verify

RUN go build -v -o ./ ./...

ENTRYPOINT [ "tail", "-f", "/dev/null" ]