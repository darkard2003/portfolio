# Build stage
FROM golang:alpine AS builder

# Allow cross-compilation by passing target OS/ARCH at build time
# Example: podman build --platform linux/arm64 --build-arg TARGETOS=linux --build-arg TARGETARCH=arm64 -t portfolio:pi .
ARG TARGETOS=linux
ARG TARGETARCH=amd64

# Set working directory
WORKDIR /app

# Install git (needed for some Go modules)
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application for the specified target
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -a -installsuffix cgo -o main .

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Create a non-root user
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .


# Change ownership to non-root user
RUN chown -R appuser:appgroup /root

# Switch to non-root user
USER appuser

# Expose port (Fly.io typically uses 8080)
EXPOSE 8080

# Run the application
CMD ["./main"]
