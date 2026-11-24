.PHONY: run-app run generate ssr-test test down

COMPOSE = docker compose -f .devcontainer/docker-compose.yml --env-file .env

# ejecuta tareas previas, levanta la aplicacion y limpia al finalizar
run-app:
	@chmod +x ./run-app.sh
	@./run-app.sh

# alias de run-app
run: run-app

# compila el acceso a datos (sqlc) y las vistas (templ)
generate:
	@echo ">> Generando codigo con sqlc y vistas con templ..."
	@$(COMPOSE) exec app sqlc generate
	@$(COMPOSE) exec app templ generate

# ejecuta las pruebas de integracion con hurl
ssr-test:
	@echo ">> Ejecutando pruebas de integracion con hurl..."
	@$(COMPOSE) exec app hurl --test ssr_requests.hurl

# ejecuta los tests unitarios sobre una base limpia
test: down
	@$(COMPOSE) up -d --build --wait
	@$(COMPOSE) exec app sqlc generate
	@$(COMPOSE) exec app templ generate
	@$(COMPOSE) exec app go test ./... -v; \
	TEST_EXIT=$$?; \
	$(COMPOSE) down -v --remove-orphans; \
	exit $$TEST_EXIT

# detiene y elimina contenedores y volumenes
down:
	@echo ">> Bajando contenedores y borrando volumenes..."
	@$(COMPOSE) down -v --remove-orphans
