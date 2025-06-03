#
# Etapa 1: compilación del Worker
#
FROM golang:1.24-alpine AS builder

# Instalamos git y certificados para que go mod pueda obtener dependencias si las hubiera
RUN apk add --no-cache git ca-certificates

# Fijamos como directorio de trabajo /app
WORKDIR /app

# 1) Copiamos únicamente go.mod y go.sum para cachear la capa de dependencias
COPY go.mod go.sum ./
RUN go mod download

# 2) Copiamos todo el código fuente del Worker
COPY . .

# 3) Compilamos el binario “worker”
RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -o worker ./cmd

#
# Etapa 2: imagen final ligera
#
FROM alpine:3.17

# Necesitamos certificados para HTTPS (si de pronto quisiéramos hacer healthcheck a HTTPS)
RUN apk add --no-cache ca-certificates

WORKDIR /root/

# Copiamos el binario compilado desde la etapa builder
COPY --from=builder /app/worker .

# Exponemos el puerto 8080 (por convención, el Worker escucha en :8080)
EXPOSE 8080

# Al arrancar el contenedor, ejecutamos el binario
ENTRYPOINT ["./worker"]
