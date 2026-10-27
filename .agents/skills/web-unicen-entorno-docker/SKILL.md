---
name: web-unicen-entorno-docker
description: >-
  Guía y comandos de la cátedra de Programación Web (UNICEN) para Docker, Docker Compose,
  Devcontainers y herramientas de desarrollo como Air (live reload) y Atlas (migraciones de BD).
---

# Entorno Docker y DevTools (Cátedra Web UNICEN)

Esta skill documenta el setup de contenedores y utilidades de desarrollo según `teoria/03-docker.md` y `teoria/06-devTools.md`.

---

## 1. Docker Compose para Base de Datos (PostgreSQL)

Configuración típica utilizada en los TPs para levantar la base de datos localmente:

```yaml
version: '3.8'

services:
  database:
    image: postgres:16-alpine
    container_name: web_db
    restart: unless-stopped
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: supersecret
      POSTGRES_DB: tp2_db
    ports:
      - "5432:5432"
    volumes:
      - pgdata:/var/lib/postgresql/data

volumes:
  pgdata:
```

### Comandos frecuentes:
- Iniciar servicios: `docker compose up -d`
- Ver logs: `docker compose logs -f database`
- Detener y limpiar: `docker compose down`

---

## 2. Dev Tools: Air (Live Reloading para Go)

En `teoria/06-devTools.md` se enseña el uso de `air` para compilar y reiniciar el servidor Go automáticamente ante cualquier cambio de código.

### Comandos:
- Ejecutar en desarrollo: `air`
- Configuración opcional en `.air.toml`:
  ```toml
  root = "."
  tmp_dir = "tmp"

  [build]
    cmd = "go build -o ./tmp/main ."
    bin = "./tmp/main"
    include_ext = ["go", "tpl", "tmpl", "html"]
    exclude_dir = ["assets", "tmp", "vendor"]
  ```

---

## 3. Dev Tools: Atlas (Gestión de Esquemas de Base de Datos)

Utilizado para inspeccionar y aplicar cambios al esquema de la base de datos de forma declarativa:
- Inspeccionar esquema actual: `atlas schema inspect -u "postgres://postgres:supersecret@localhost:5432/tp2_db?sslmode=disable"`
- Aplicar esquema desde archivo SQL: `atlas schema apply -u "postgres://postgres:supersecret@localhost:5432/tp2_db?sslmode=disable" --to "file://db/schema/users.sql"`

---

## 4. Dev Container Unificado del Repositorio

Para no duplicar contenedores por cada TP, se utiliza un único entorno en `.devcontainer/`:
- **Montaje**: Monta la raíz en `/workspace`, permitiendo trabajar en `tp1/`, `tp2/`, `tp3/` y ver `teoria/` sin cambiar de ventana.
- **Herramientas incluidas**: Go 1.27, `sqlc`, `hurl`, `postgresql-client`.
- **Base de Datos**: Servicio `database` con PostgreSQL 15 y healthcheck activo en `database:5432`.


