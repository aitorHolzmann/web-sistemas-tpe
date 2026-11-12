#!/usr/bin/env bash
set -e

# Asegurar que se ejecute desde la raíz del proyecto (directorio del script)
cd "$(dirname "$0")"

COMPOSE="docker compose -f .devcontainer/docker-compose.yml --env-file .env"

# 3) TAREAS POSTERIORES: borrar contenedores y volúmenes
cleanup() {
    echo ""
    echo "================================================================="
    echo ">> [3/3] TAREAS POSTERIORES: Borrando contenedores y volúmenes..."
    echo "================================================================="
    $COMPOSE down -v --remove-orphans || true
}
# Atrapa la salida normal o interrupción (Ctrl+C, SIGTERM) para garantizar la limpieza posterior
trap cleanup EXIT INT TERM

echo "================================================================="
echo ">> [1/3] TAREAS PREVIAS"
echo "================================================================="

echo "-> 1. Limpieza inicial: Borrando contenedores y volúmenes previos..."
$COMPOSE down -v --remove-orphans

echo "-> 2. Levantando contenedores (construyendo imágenes si es necesario)..."
$COMPOSE up -d --build --wait

echo "-> 3. Esperando que la base de datos esté lista y respondiendo..."
$COMPOSE exec db sh -c 'until pg_isready -U "$$POSTGRES_USER" -d "$$POSTGRES_DB"; do sleep 1; done'

echo "-> 4. Ejecutando generación de acceso a datos con sqlc..."
$COMPOSE exec app sqlc generate

echo "-> 5. Compilando componentes de vista con templ..."
$COMPOSE exec app templ generate

echo "-> 6. Compilando la aplicación Go..."
$COMPOSE exec app go build -o /tmp/server main.go

echo "================================================================="
echo ">> [2/3] EJECUTANDO LA APLICACIÓN"
echo "================================================================="
echo "Servidor disponible en: http://localhost:8080"
echo "Presione Ctrl+C para detener el servidor y ejecutar automáticamente las tareas posteriores."
echo "================================================================="

# Ejecutamos la aplicación
$COMPOSE exec app /tmp/server
