#!/usr/bin/env bash
set -e

# ejecuta desde la raiz del proyecto
cd "$(dirname "$0")"

COMPOSE="docker compose -f .devcontainer/docker-compose.yml --env-file .env"

# tareas posteriores: baja contenedores y borra volumenes al salir
cleanup() {
    echo ""
    echo ">> [3/3] Tareas posteriores: borrando contenedores y volumenes..."
    $COMPOSE down -v --remove-orphans || true
}
trap cleanup EXIT

echo ">> [1/3] Tareas previas"

echo "-> 1. Limpiando contenedores y volumenes previos..."
$COMPOSE down -v --remove-orphans

echo "-> 2. Levantando contenedores..."
$COMPOSE up -d --build --wait

echo "-> 3. Esperando que la base de datos este lista..."
$COMPOSE exec db sh -c 'until pg_isready -U "$POSTGRES_USER" -d "$POSTGRES_DB"; do sleep 1; done'

echo "-> 4. Generando codigo con sqlc..."
$COMPOSE exec app sqlc generate

echo "-> 5. Compilando vistas con templ..."
$COMPOSE exec app templ generate

echo "-> 6. Compilando la aplicacion..."
$COMPOSE exec app go build -o /tmp/server main.go

echo ">> [2/3] Ejecutando la aplicacion en http://localhost:8080 (Ctrl+C para detener)"
$COMPOSE exec app /tmp/server
