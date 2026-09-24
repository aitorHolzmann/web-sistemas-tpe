.PHONY: test sqlc build up down run api-test

COMPOSE = docker compose -f .devcontainer/docker-compose.yml --env-file .env

test: sqlc build down up
	@echo ""
	@echo ">> Corriendo tests..."
	@$(COMPOSE) run --rm app go test ./... -v; \
	TEST_EXIT=$$?; \
	echo ""; \
	echo ">> Bajando contenedores y volúmenes..."; \
	$(COMPOSE) down -v; \
	exit $$TEST_EXIT

sqlc:
	@echo ">> Generando código con sqlc..."
	@$(COMPOSE) run --rm app sqlc generate

build:
	@echo ">> Descargando dependencias y compilando..."
	@$(COMPOSE) run --rm app sh -c "go mod tidy && go build -buildvcs=false ./..."

up:
	@echo ">> Levantando la base de datos..."
	@$(COMPOSE) up -d --wait db

down:
	@echo ">> Bajando contenedores y borrando volúmenes..."
	@$(COMPOSE) down -v --remove-orphans

run: sqlc build up
	@echo ">> Iniciando servidor en puerto 8080..."
	@$(COMPOSE) run --rm --service-ports app go run main.go

api-test:
	@echo ">> Ejecutando pruebas de API..."
	@./requests.sh
