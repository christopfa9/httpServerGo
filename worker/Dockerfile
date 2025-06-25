#
# Stage 1: build the Worker
#
FROM golang:1.24-alpine AS builder

# Install git and certificates so that go mod can fetch dependencies if needed
RUN apk add --no-cache git ca-certificates

# Set /app as the working directory
WORKDIR /app

# 1) Copy only go.mod and go.sum to cache the dependency layer
COPY go.mod go.sum ./
RUN go mod download

# 2) Copy all Worker source code
COPY . .

# 3) Compile the “worker” binary
RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -o worker ./cmd

#
# Stage 2: final lightweight image
#
FROM alpine:3.17

# Certificates are needed for HTTPS (e.g., if we want to health check via HTTPS)
RUN apk add --no-cache ca-certificates

WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/worker .

# Expose port 8080 (by convention, the Worker listens on :8080)
EXPOSE 8080

# When the container starts, run the binary
ENTRYPOINT ["./worker"]
