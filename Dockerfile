# Build stage
FROM golang:1.23-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o elotus .

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Create a non-root user
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

# Set working directory
WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/elotus .

# Copy config file
COPY --from=builder /app/config.yaml .

# Change ownership to non-root user
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose ports (will be overridden in docker-compose)
EXPOSE 8080 8081

# Default command (will be overridden in docker-compose)
CMD ["./elotus"] 