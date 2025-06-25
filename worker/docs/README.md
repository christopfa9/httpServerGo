TEC

---

# HTTP Server Go

**HTTP Server Go** is a sample project in Go that implements a minimal HTTP server with multiple endpoints for various operations (calculation, text manipulation, file operations, task simulation, metrics, etc.). It also includes a _worker pool_ system to horizontally scale the execution of commands.

---

## Features

- **Diverse endpoints**:

  - `/help` – Lists all available endpoints and their parameters
  - `/fibonacci` – Computes the n-th Fibonacci number
  - `/createfile` – Creates or truncates a file, writing repeated content
  - `/deletefile` – Deletes an existing file
  - `/reverse` – Reverses a UTF-8-safe string
  - `/toupper` – Converts text to uppercase
  - `/random` – Generates an array of random numbers
  - `/timestamp` – Returns the current time in ISO-8601 format
  - `/hash` – Computes the SHA-256 hash of a text
  - `/simulate` – Simulates a task by sleeping for X seconds (with optional name)
  - `/sleep` – Pauses execution for X seconds
  - `/loadtest` – Runs N concurrent tasks sleeping S seconds each
  - `/status` – Shows server metrics (uptime, connections, processes)

- **Worker pools**: Each heavy command (Fibonacci, loadtest, simulate, etc.) is handled by a pool of dedicated goroutines, allowing load balancing and serving multiple requests concurrently without blocking the main listener.

- **Docker ready**: Containerized with multi-stage builds to generate a static and lightweight binary.

---

## Installation and Execution

### 1. Clone the repository

```bash
git clone https://github.com/YOUR_USERNAME/YOUR_REPO.git
cd YOUR_REPO
```

### 2. Compile locally

```bash
go mod tidy
go build -o httpserver ./cmd
./httpserver
```

The server starts on port `8080` by default.

### 3. Using Docker

1. **Build the image**:

   ```bash
   docker build -t httpservergo:latest .
   ```

2. **Run the container**:

   ```bash
   docker run -d --name httpsrv -p 8080:8080 httpservergo:latest
   ```

3. **View logs**:

   ```bash
   docker logs -f httpsrv
   ```

---

## Usage Examples

```bash
curl "http://localhost:8080/help"
curl "http://localhost:8080/fibonacci?num=10"
curl "http://localhost:8080/createfile?name=foo.txt&content=Hola&repeat=3"
# ...more endpoints as listed
curl "http://localhost:8080/status"
```

---

## Internal Architecture

- `cmd/main.go` – Entry point, initializes pools and listener
- `internal/server/` – HTTP listening and handler logic

  - `listener.go` – Accepts connections and dispatches them
  - `handler.go` – Parses requests and queues them into pools
  - `workerPools.go` – Worker pool setup and configuration

- `internal/commands/` – Implementation of each command
- `internal/status/` – Server metrics (uptime, connections, processes)
- `internal/utils/` – Helper functions (parsing, HTTP responses, sanitization)

---

## Contributing

Contributions are welcome. Please open an issue or submit a pull request with new features or improvements.

---

## License

This project is for educational use and does not include an explicit license.

---
