# Etapa 1: build
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copiamos go.mod y go.sum primero (para aprovechar cache de dependencias)
COPY go.mod go.sum ./
RUN go mod download

# Copiamos el resto del código
COPY . .

# Generar la documentación Swagger (por si no está actualizada)
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init --parseDependency --parseInternal

# Compilar el binario
RUN go build -o main .

# Etapa 2: runtime
FROM alpine:3.20

WORKDIR /app

# Copiar binario desde la etapa anterior
COPY --from=builder /app/main .
COPY .env .env
COPY docs ./docs

# Exponer el puerto del servidor
EXPOSE 8080

# Ejecutar la API
CMD ["./main"]
