# Stage 1: Build Stage
FROM golang:1.23.2-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first for caching dependencies
COPY go.mod go.sum ./

RUN go mod download

# Copy the entire application code
COPY . .

# Build the application binary
RUN go build -o main .

# Stage 2: Final Stage
FROM gcr.io/distroless/base-debian10

WORKDIR /app

# Copy the compiled binary from the build stage
COPY --from=builder /app/main /app/main

# Copy additional files
COPY ./docs /app/docs
COPY .env /app/.env

# Expose port for the application
EXPOSE 8000

# Command to run the application
CMD ["/app/main"]
