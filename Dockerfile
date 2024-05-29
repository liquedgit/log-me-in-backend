# Dockerfile
FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN until go mod tidy; do echo "Retrying..."; done
RUN go build -o main .

COPY mysql-healthcheck.sh /usr/local/bin/mysql-healthcheck.sh
RUN chmod +x /usr/local/bin/mysql-healthcheck.sh

EXPOSE 1234

CMD ["./main"]
