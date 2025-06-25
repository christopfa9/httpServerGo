---
# How to Run the Program with Docker

This document describes the steps to build the Docker image, run the server container, and test the endpoints.
---

## 1) Build the Docker Image

From the project root (where the Dockerfile is located):

```bash
docker build -t httpservergo:latest .
```

- `-t httpservergo:latest` tags the image.
- The dot `.` sets the current directory as the build context.

---

## 2) Run the Container

```bash
docker run -d --name httpservergo_container -p 8080:8080 httpservergo:latest
```

- `-d` runs the container in detached mode.
- `--name` assigns a name to the container.
- `-p 8080:8080` maps container port 8080 to host port 8080.

---

## 3) View Server Logs

```bash
docker logs -f httpservergo_container
```

---

## 4) Test Endpoints with curl (or a browser)

- **Help / list of commands**

  ```bash
  curl "http://localhost:8080/help"
  ```

- **Fibonacci (10th number)**

  ```bash
  curl "http://localhost:8080/fibonacci?num=10"
  ```

- **Create a file**

  ```bash
  curl "http://localhost:8080/createfile?name=test.txt&content=Hola&repeat=3"
  ```

- **Delete a file**

  ```bash
  curl "http://localhost:8080/deletefile?name=test.txt"
  ```

- **Reverse text**

  ```bash
  curl "http://localhost:8080/reverse?text=HolaMundo"
  ```

- **Convert to uppercase**

  ```bash
  curl "http://localhost:8080/toupper?text=hola-mundo"
  ```

- **Random numbers (5 between 1 and 100)**

  ```bash
  curl "http://localhost:8080/random?count=5&min=1&max=100"
  ```

- **Timestamp (current time)**

  ```bash
  curl "http://localhost:8080/timestamp"
  ```

- **SHA-256 hash of text**

  ```bash
  curl "http://localhost:8080/hash?text=datossecretos"
  ```

- **Simulate task (4s, optional name)**

  ```bash
  curl "http://localhost:8080/simulate?seconds=4&task=demoTask"
  ```

- **Sleep (2s pause)**

  ```bash
  curl "http://localhost:8080/sleep?seconds=2"
  ```

- **Load test (10 concurrent tasks sleeping 1s each)**

  ```bash
  curl "http://localhost:8080/loadtest?tasks=10&sleep=1"
  ```

- **Server metrics**

  ```bash
  curl "http://localhost:8080/status"
  ```

---

## 5) Stop and Remove the Container

To stop:

```bash
docker stop httpservergo_container
```

To remove:

```bash
docker rm httpservergo_container
```

---

## 6) Rebuild the Image After Changes

If you update the code or Dockerfile:

```bash
docker build -t httpservergo:latest .
docker stop httpservergo_container
docker rm httpservergo_container
docker run -d --name httpservergo_container -p 8080:8080 httpservergo:latest
```

---
