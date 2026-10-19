.PHONY: run api-test test down

COMPOSE = docker compose -f .devcontainer/docker-compose.yml --env-file .env

# levanta los contenedores, genera el codigo de sqlc e inicia el servidor en el puerto 8080
run:
	@echo ">> Levantando contenedores (base de datos y app)..."
	@$(COMPOSE) up -d --build --wait
	@echo ">> Generando código con sqlc..."
	@$(COMPOSE) exec app sqlc generate
	@echo ">> Iniciando servidor en http://localhost:8080 ..."
	@$(COMPOSE) exec app go run main.go

# ejecuta las pruebas de la API con hurl (el servidor tiene que estar corriendo)
api-test:
	@echo ">> Ejecutando pruebas de API..."
	@$(COMPOSE) exec app hurl --test requests.hurl

# corre los tests de la capa de datos (sqlc) sobre una base de datos limpia
test: down
	@$(COMPOSE) up -d --build --wait
	@$(COMPOSE) exec app sqlc generate
	@$(COMPOSE) exec app go test ./... -v; \
	TEST_EXIT=$$?; \
	$(COMPOSE) down -v; \
	exit $$TEST_EXIT

# baja los contenedores y borra los datos de la base
down:
	@echo ">> Bajando contenedores y borrando volúmenes..."
	@$(COMPOSE) down -v --remove-orphans
