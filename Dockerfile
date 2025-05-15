# Etapa 1: compilación
FROM golang:1.24-alpine AS builder

# Instala dependencias OS necesarias (cURL, certificados, etc.)
RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Copia módulos y cachea
COPY go.mod go.sum ./
RUN go mod download

# Copia el código fuente
COPY . .

# Compila en modo release
RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -o httpserver ./cmd

# Etapa 2: imagen final
FROM alpine:3.17

# Permite verificar HTTPS si usas TLS, etc.
RUN apk add --no-cache ca-certificates

WORKDIR /root/

# Copia sólo el binario estático
COPY --from=builder /app/httpserver .

# Exponer el puerto en el que corre tu servidor
EXPOSE 8080

# Variable de entorno (opcional, coincide con tu main.go)
ENV PORT=8080

# Comando por defecto
ENTRYPOINT ["./httpserver"]