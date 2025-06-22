FROM golang:1.22-alpine as builder

WORKDIR /app
COPY backend/ ./backend/
COPY go.mod ./
RUN go build -o server ./backend/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/server .
COPY frontend/ ./frontend/
EXPOSE 8080
CMD ["./server"]
