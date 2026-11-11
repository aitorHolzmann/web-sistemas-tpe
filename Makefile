.PHONY: run generate ssr-test test down

COMPOSE = docker compose -f .devcontainer/docker-compose.yml --env-file .env

# levanta los contenedores, genera sqlc y templ e inicia el servidor en el puerto 8080
run:
	@echo ">> Levantando contenedores (base de datos y app)..."
	@$(COMPOSE) up -d --build --wait
	@echo ">> Generando código con sqlc y vistas con templ..."
	@$(COMPOSE) exec app sqlc generate
	@$(COMPOSE) exec app templ generate
	@echo ">> Iniciando servidor SSR en http://localhost:8080 ..."
	@$(COMPOSE) exec app go run main.go

# compila el acceso a datos (sqlc) y los componentes de vista (templ)
generate:
	@echo ">> Generando código con sqlc y vistas con templ..."
	@$(COMPOSE) exec app sqlc generate
	@$(COMPOSE) exec app templ generate

# ejecuta las pruebas de integración SSR y PRG con hurl (el servidor debe estar corriendo)
ssr-test:
	@echo ">> Ejecutando pruebas de integración SSR con hurl..."
	@$(COMPOSE) exec app hurl --test ssr_requests.hurl

# corre los tests de la capa de datos (sqlc) sobre una base de datos limpia
test: down
	@$(COMPOSE) up -d --build --wait
	@$(COMPOSE) exec app sqlc generate
	@$(COMPOSE) exec app templ generate
	@$(COMPOSE) exec app go test ./... -v; \
	TEST_EXIT=$$?; \
	$(COMPOSE) down -v; \
	exit $$TEST_EXIT

# baja los contenedores y borra los datos de la base
down:
	@echo ">> Bajando contenedores y borrando volúmenes..."
	@$(COMPOSE) down -v --remove-orphans
