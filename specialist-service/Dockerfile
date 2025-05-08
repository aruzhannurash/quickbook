# syntax=docker/dockerfile:1
FROM golang:1.23


WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o specialist-service .

EXPOSE 8082

CMD ["./specialist-service"]
