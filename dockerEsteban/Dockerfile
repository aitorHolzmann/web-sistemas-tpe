# ------------------------------------------------------------------
# Dockerfile OPCIONAL para esta etapa del TP2 (Trabajo de Cursada).
#
# La consigna es explícita: "no es necesario conectar la capa de datos
# a un servidor web" todavía. Por eso este Dockerfile NO levanta ningún
# servicio; solo sirve para comprobar, en un entorno reproducible, que
# el módulo Go junto con el código generado por sqlc (./db/sqlc)
# compila sin errores. Es útil como paso de CI o como sanity-check
# antes de entregar.
# ------------------------------------------------------------------

# Etapa 1: compilación / verificación
FROM golang:1.23-alpine AS builder
WORKDIR /app

# Copiamos primero los manifiestos para aprovechar la caché de capas:
# si go.mod/go.sum no cambian, Docker reutiliza la capa de dependencias.
COPY go.mod go.sum* ./
RUN go mod download

# Copiamos el resto del código (incluye db/sqlc, generado previamente
# con `sqlc generate` fuera o dentro del contenedor).
COPY . .

# go build sin -o de un binario final: solo verificamos que TODO el
# módulo (incluido el paquete db generado) compila correctamente.
RUN go build ./...
RUN go vet ./...

# Etapa 2: imagen mínima, solo deja constancia de que el build fue exitoso.
# No es un servicio de larga duración: no hay EXPOSE ni servidor HTTP.
FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app /app
CMD ["echo", "Paquete de acceso a datos compilado correctamente. Aun sin capa web (por diseño de esta etapa)."]
