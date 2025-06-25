#
# Stage 1: build the Dispatcher
#
FROM golang:1.24-alpine AS builder

# 1) Install git and certificates so that go mod download works
RUN apk add --no-cache git ca-certificates

# 2) Set the working directory
WORKDIR /app

# 3) Copy only go.mod and go.sum to cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# 4) Copy all source code of the Dispatcher
COPY . .

# 5) Compile the “dispatcher” binary pointing to ./cmd
RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -o dispatcher ./cmd

#
# Stage 2: final lightweight image
#
FROM alpine:3.17

# 6) Install certificates (for TLS/HTTPS if needed)
RUN apk add --no-cache ca-certificates

# 7) Working directory in the final image
WORKDIR /root/

# 8) Copy the binary from the builder stage
COPY --from=builder /app/dispatcher .

# 9) Expose port 9090 (where the Dispatcher listens)
EXPOSE 9090

# 10) Entry point to run the Dispatcher when the container starts
ENTRYPOINT ["./dispatcher"]
