FROM golang:1.24-alpine

# Install git (needed for go get) and Air
RUN apk add --no-cache git curl
RUN go install github.com/air-verse/air@latest

WORKDIR /app
COPY .air.toml ./
# Copy go.mod first for better caching
COPY go.mod ./
# (If you have go.sum, add: COPY go.sum ./)
RUN go mod download

# Copy the rest of the code
COPY backend/ ./backend/
COPY frontend/ ./frontend/

EXPOSE 8080

# Use Air for live reload (default config)
CMD ["air"]
