FROM golang:1.24-alpine

RUN apk add --no-cache git curl
RUN go install github.com/air-verse/air@latest

WORKDIR /app
COPY .air.toml ./
COPY go.mod ./
RUN go mod download

COPY backend/ ./backend/
COPY frontend/ ./frontend/

EXPOSE 8080

CMD ["air"]
