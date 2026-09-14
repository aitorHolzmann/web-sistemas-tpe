# ---------------------------------------------------------------------------
# Makefile del TP2.
#
# Si nunca usaste "make": es un ejecutor de tareas. Cada bloque de abajo con
# forma "nombre:" es un "target" (una tarea). Se ejecuta con "make nombre".
# Las líneas indentadas debajo son los comandos de shell que corre esa tarea,
# en orden, uno por uno. ".PHONY" simplemente le avisa a make que estos
# nombres son tareas, no archivos que hay que buscar en el disco.
#
# Todo corre adentro de contenedores Docker: en la máquina donde se evalúa
# este TP solo hace falta tener instalados Docker y make. No se necesita
# tener Go ni sqlc instalados en el host.
# ---------------------------------------------------------------------------

.PHONY: test sqlc build up down

# COMPOSE es un atajo: en vez de repetir esta línea larga en cada target,
# la definimos una vez y la reusamos como $(COMPOSE).
#   -f                  -> le decimos a qué docker-compose.yml apuntar
#                          (vive en .devcontainer/, no en la raíz)
#   --project-directory . -> resuelve el .env desde la raíz del repo,
#                          no desde .devcontainer/
COMPOSE = docker compose -f .devcontainer/docker-compose.yml --project-directory . --env-file .env

# "make test" (o directamente "make", ver .DEFAULT_GOAL abajo) es el punto
# de entrada único que pide la consigna. Hace, en orden:
#   1) tareas previas: sqlc generate, compilar, bajar y volver a levantar
#      los contenedores desde cero, esperar a que la base esté sana
#   2) correr los tests
#   3) tareas posteriores: bajar contenedores y volúmenes
#
# El "@" al principio de una línea le dice a make que no imprima el comando
# en sí antes de ejecutarlo (para que la salida sea más limpia).
test: sqlc build down up
	@echo ""
	@echo ">> Corriendo tests..."
	@$(COMPOSE) run --rm app go test ./... -v; \
	TEST_EXIT=$$?; \
	echo ""; \
	echo ">> Bajando contenedores y volúmenes..."; \
	$(COMPOSE) down -v; \
	exit $$TEST_EXIT

# Genera el paquete db/sqlc a partir de db/schema y db/queries (ver sqlc.yaml).
# Se corre DENTRO del contenedor "app", que ya trae el binario de sqlc
# instalado (ver .devcontainer/Dockerfile). No hace falta que sqlc esté
# instalado en tu máquina.
sqlc:
	@echo ">> Generando código con sqlc..."
	@$(COMPOSE) run --rm app sqlc generate

# Compila todo el proyecto (incluye el paquete generado por sqlc y los
# tests, porque "go build" también verifica que los _test.go compilen,
# aunque no los ejecuta). "go mod tidy" descarga las dependencias del
# go.mod y genera/actualiza go.sum si hiciera falta.
build:
	@echo ">> Descargando dependencias y compilando..."
	@$(COMPOSE) run --rm app sh -c "go mod tidy && go build ./..."

# Levanta el contenedor de la base y espera a que su healthcheck pase
# (ver "healthcheck" en docker-compose.yml). "--wait" es justamente la
# forma que tiene Docker Compose de decir "no me devuelvas el control
# hasta que el servicio esté healthy".
up:
	@echo ">> Levantando la base de datos..."
	@$(COMPOSE) up -d --wait db

# Apaga y borra los contenedores Y los volúmenes (-v). Borrar el volumen
# es clave: así cada corrida arranca con una base 100% limpia, con el
# schema.sql aplicado desde cero.
down:
	@echo ">> Bajando contenedores y borrando volúmenes..."
	@$(COMPOSE) down -v --remove-orphans
