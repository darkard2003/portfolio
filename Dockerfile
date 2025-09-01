# Build stage
FROM golang:1.24-alpine AS builder

# Install Node.js for TailwindCSS and templ for template generation
RUN apk add --no-cache nodejs npm
RUN go install github.com/a-h/templ/cmd/templ@latest

WORKDIR /app

# Copy package files and install dependencies
COPY package*.json ./
RUN npm ci
RUN npm install -g @tailwindcss/cli

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Generate templates
RUN go generate ./...

# Generate CSS
RUN npx @tailwindcss/cli -i ./web/css/input.css -o ./static/css/app.css --minify

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/web/main.go

# Production stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Copy static assets
COPY --from=builder /app/static ./static

# Copy data file
COPY --from=builder /app/data.json .

# Expose port
EXPOSE 8080

# Set environment variable for production
ENV PORT=8080

# Run the binary
CMD ["./main"]