.PHONY: run-app run generate ssr-test test down

COMPOSE = docker compose -f .devcontainer/docker-compose.yml --env-file .env

# Script completo exigido por las consideraciones de entrega (tareas previas, ejecución y limpieza posterior)
run-app:
	@chmod +x ./run-app.sh
	@./run-app.sh

# Inicia la aplicación ejecutando las tareas previas, levantando el servidor y limpiando al finalizar
run: run-app

# Compila el acceso a datos (sqlc) y los componentes de vista (templ)
generate:
	@echo ">> Generando código con sqlc y vistas con templ..."
	@$(COMPOSE) exec app sqlc generate
	@$(COMPOSE) exec app templ generate

# Ejecuta las pruebas de integración SSR y PRG con hurl (el servidor debe estar corriendo)
ssr-test:
	@echo ">> Ejecutando pruebas de integración SSR con hurl..."
	@$(COMPOSE) exec app hurl --test ssr_requests.hurl

# Corre los tests de la capa de datos (sqlc) sobre una base de datos limpia
test: down
	@$(COMPOSE) up -d --build --wait
	@$(COMPOSE) exec app sqlc generate
	@$(COMPOSE) exec app templ generate
	@$(COMPOSE) exec app go test ./... -v; \
	TEST_EXIT=$$?; \
	$(COMPOSE) down -v --remove-orphans; \
	exit $$TEST_EXIT

# Baja los contenedores y borra los datos de la base
down:
	@echo ">> Bajando contenedores y borrando volúmenes..."
	@$(COMPOSE) down -v --remove-orphans
