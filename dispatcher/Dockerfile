#
# Etapa 1: compilación del Dispatcher
#
FROM golang:1.24-alpine AS builder

# 1) Instalamos git y certificados para que go mod download funcione
RUN apk add --no-cache git ca-certificates

# 2) Establecemos el directorio de trabajo
WORKDIR /app

# 3) Copiamos solo go.mod y go.sum para cachear dependencias
COPY go.mod go.sum ./
RUN go mod download

# 4) Copiamos todo el código fuente del Dispatcher
COPY . .

# 5) Compilamos el binario “dispatcher” apuntando a ./cmd
RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -o dispatcher ./cmd

#
# Etapa 2: imagen final ligera
#
FROM alpine:3.17

# 6) Instalamos certificados (por si se necesita TLS/HTTPS)
RUN apk add --no-cache ca-certificates

# 7) Directorio de trabajo en imagen final
WORKDIR /root/

# 8) Copiamos el binario desde la etapa builder
COPY --from=builder /app/dispatcher .

# 9) Exponemos el puerto 9090 (donde escucha el Dispatcher)
EXPOSE 9090

# 10) Entry point para ejecutar el Dispatcher al arrancar el contenedor
ENTRYPOINT ["./dispatcher"]
