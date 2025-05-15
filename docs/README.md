# HTTP Server Go

**HTTP Server Go** es un proyecto de ejemplo en Go que implementa un servidor HTTP minimalista con múltiples endpoints para diferentes operaciones (cálculo, manipulación de texto, operaciones en archivos, simulación de tareas, métricas, etc.). Además, incluye un sistema de *worker pools* para escalar horizontalmente la ejecución de comandos.

---

## Características

* **Endpoints diversos**:

  * `/help`       – Lista todos los endpoints disponibles con sus parámetros.
  * `/fibonacci`  – Calcula el n-ésimo número de Fibonacci.
  * `/createfile` – Crea o trunca un archivo, escribiendo contenido repetido.
  * `/deletefile` – Elimina un archivo existente.
  * `/reverse`    – Invierte una cadena de texto (UTF-8 seguro).
  * `/toupper`    – Convierte texto a mayúsculas.
  * `/random`     – Genera un arreglo de números aleatorios.
  * `/timestamp`  – Devuelve la hora actual en formato ISO-8601.
  * `/hash`       – Calcula el hash SHA-256 de un texto.
  * `/simulate`   – Simula una tarea durmiendo X segundos (con nombre opcional).
  * `/sleep`      – Pausa la ejecución durante X segundos.
  * `/loadtest`   – Ejecuta N tareas concurrentes pausando S segundos cada una.
  * `/status`     – Muestra métricas del servidor (uptime, conexiones, procesos).

* **Worker pools**: Cada comando pesado (Fibonacci, loadtest, simulate, etc.) es atendido por un pool de goroutines dedicadas, permitiendo balancear la carga y atender múltiples peticiones simultáneas sin bloquear el *acceptor*.

* **Docker ready**: Contenerizado con multi-stage build para generar un binario estático y ligero.

---

## Instalación y ejecución

### 1. Clonar el repositorio

```bash
git clone https://github.com/TU_USUARIO/TU_REPO.git
cd TU_REPO
```

### 2. Compilar localmente

```bash
go mod tidy
go build -o httpserver ./cmd
./httpserver
```

El servidor arranca en el puerto `8080` por defecto.

### 3. Usando Docker

1. **Construir imagen**:

   ```bash
   ```

docker build -t httpservergo\:latest .

````
2. **Ejecutar contenedor**:
   ```bash
docker run -d --name httpsrv -p 8080:8080 httpservergo:latest
````

3. **Ver logs**:

   ```bash
   ```

docker logs -f httpsrv

````

---

## Ejemplos de uso

```bash
curl "http://localhost:8080/help"
curl "http://localhost:8080/fibonacci?num=10"
curl "http://localhost:8080/createfile?name=foo.txt&content=Hola&repeat=3"
# ...otros endpoints según lista
curl "http://localhost:8080/status"
````

---

## Arquitectura interna

* `cmd/main.go`        – Punto de entrada, inicializa pools y listener.
* `internal/server/`   – Lógica de escucha y *handler* HTTP.

  * `listener.go`      • Acepta conexiones y las despacha.
  * `handler.go`       • Parseo de requests y encolado en pools.
  * `workerPools.go`   • Definición e inicialización de worker pools.
* `internal/commands/` – Implementación de cada comando.
* `internal/status/`   – Métricas del servidor (*uptime*, conexiones, procesos).
* `internal/utils/`    – Funciones auxiliares (parsing, respuestas HTTP, sanitización).

---

## Contribuciones

Las contribuciones son bienvenidas. Por favor, abre un *issue* o envía un *pull request* con nuevas funcionalidades o mejoras.

---

## Licencia

Este proyecto es de uso educativo y no tiene licencia explícita.").
