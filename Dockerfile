# Build stage
FROM golang:latest AS builder

WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /goserver

EXPOSE 8080

CMD ["/app/surely"]