# Etapa 1: Construcción
FROM golang:1.20-alpine AS builder

WORKDIR /app

# Copia los archivos de dependencias y descarga los módulos
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copia el código fuente y compila el binario
COPY . .
RUN go build -o main .

# Etapa 2: Imagen final en entorno más ligero
FROM alpine:latest

WORKDIR /root/

# Copia el binario desde la etapa de construcción
COPY --from=builder /app/main .

# Instalar las dependencias necesarias para ejecutar el binario (si es necesario)
RUN apk --no-cache add ca-certificates

# Exponer el puerto
EXPOSE 8000

# Dockerize para esperar a la db, idealmente la db está en una instancia remota siempre prendida, no need de esto
RUN apk add --no-cache curl \
    && curl -sSL https://github.com/jwilder/dockerize/releases/download/v0.6.1/dockerize-alpine-linux-amd64-v0.6.1.tar.gz | tar -xz -C /usr/local/bin

# Comando de inicio modificado en el Dockerfile
CMD ["dockerize", "-wait", "tcp://db:5432", "-timeout", "45s", "./main"]

