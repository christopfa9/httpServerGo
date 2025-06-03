# Cómo correr el programa con Docker

Este documento describe los pasos para compilar la imagen Docker y levantar el servidor, así como probar los endpoints.

---

## 1) Construir la imagen Docker

Desde la raíz del proyecto (donde está el Dockerfile):

```bash
docker build -t httpservergo:latest .
```

- `-t httpservergo:latest` etiqueta la imagen.
- El punto `.` indica el contexto actual.

---

## 2) Ejecutar el contenedor

```bash
docker run -d --name httpservergo_container -p 8080:8080 httpservergo:latest
```

- `-d` ejecuta en segundo plano.
- `--name` asigna un nombre al contenedor.
- `-p 8080:8080` expone el puerto 8080.

---

## 3) Ver logs del servidor

```bash
docker logs -f httpservergo_container
```

---

## 4) Probar endpoints con curl (o navegador)

- **Ayuda / listado de comandos**  
  ```
  curl "http://localhost:8080/help"
  ```

- **Fibonacci (10º número)**  
  ```
  curl "http://localhost:8080/fibonacci?num=10"
  ```

- **Crear archivo**  
  ```
  curl "http://localhost:8080/createfile?name=test.txt&content=Hola&repeat=3"
  ```

- **Eliminar archivo**  
  ```
  curl "http://localhost:8080/deletefile?name=test.txt"
  ```

- **Invertir texto**  
  ```
  curl "http://localhost:8080/reverse?text=HolaMundo"
  ```

- **Convertir a mayúsculas**  
  ```
  curl "http://localhost:8080/toupper?text=hola-mundo"
  ```

- **Números aleatorios (5 números entre 1 y 100)**  
  ```
  curl "http://localhost:8080/random?count=5&min=1&max=100"
  ```

- **Timestamp (hora actual)**  
  ```
  curl "http://localhost:8080/timestamp"
  ```

- **SHA-256 de un texto**  
  ```
  curl "http://localhost:8080/hash?text=datossecretos"
  ```

- **Simular tarea (4s, nombre opcional)**  
  ```
  curl "http://localhost:8080/simulate?seconds=4&task=demoTask"
  ```

- **Sleep (pausa de 2s)**  
  ```
  curl "http://localhost:8080/sleep?seconds=2"
  ```

- **Load test (10 tareas concurrentes durmiendo 1s cada una)**  
  ```
  curl "http://localhost:8080/loadtest?tasks=10&sleep=1"
  ```

- **Métricas del servidor**  
  ```
  curl "http://localhost:8080/status"
  ```

---

## 5) Detener y eliminar el contenedor

Para detener:
```bash
docker stop httpservergo_container
```

Para eliminar:
```bash
docker rm httpservergo_container
```

---

## 6) Reconstruir imagen tras cambios

Si modificas código o Dockerfile:
```bash
docker build -t httpservergo:latest .
docker stop httpservergo_container
docker rm httpservergo_container
docker run -d --name httpservergo_container -p 8080:8080 httpservergo:latest
```