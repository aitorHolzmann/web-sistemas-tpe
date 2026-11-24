# Trabajo Especial - TP6: Interfaces Dinámicas con HTMX

Aplicación web de **MarkeTandil** para gestión de productos y categorías, desarrollada en Go con **Templ** (Server-Side Rendering), **Pico.css** y **HTMX 4.0** para actualizaciones parciales del DOM sin recargar la página completa.

El proyecto no incluye archivos JavaScript (`.js`) propios; la interactividad se maneja mediante atributos declarativos de HTMX cargado por CDN en el layout.

---

## Requisitos

- **Docker** (con Docker Compose)
- **Make** o **Bash**

Go, `sqlc`, `templ`, PostgreSQL y Hurl corren dentro de los contenedores definidos en `.devcontainer`.

---

## Ejecución

Desde la raíz del proyecto ejecutar:

```bash
make run-app
```

O directamente con el script bash:

```bash
./run-app.sh
```

El script automatiza las tres etapas pedidas:

1. **Tareas previas**:
   - Elimina contenedores y volúmenes previos (`docker compose down -v --remove-orphans`).
   - Construye y levanta los contenedores (`db` y `app`).
   - Espera a que PostgreSQL acepte conexiones con `pg_isready`.
   - Genera el código de acceso a datos con `sqlc generate`.
   - Compila las vistas `.templ` con `templ generate`.
   - Compila el binario del servidor Go.
2. **Ejecución de la aplicación**:
   - Levanta el servidor en **http://localhost:8080**.
3. **Tareas posteriores**:
   - Al cortar el servidor con `Ctrl+C` (o al finalizar el proceso), el `trap EXIT` ejecuta automáticamente `docker compose down -v --remove-orphans` para borrar los contenedores y volúmenes.

---

## Pruebas

### Pruebas de integración con Hurl

Con la aplicación corriendo, en otra terminal ejecutar:

```bash
make ssr-test
```

Ejecuta `ssr_requests.hurl` validando:
- Carga inicial de `GET /` (estado vacío, script de HTMX 4.0 y atributos `hx-post`, `hx-target`, `hx-swap`).
- Alta sin recarga con `POST /productos` (devuelve HTTP 200 con el fragmento actualizado `#lista-productos` y `<hx-partial>`).
- Consulta del fragmento de lista en `GET /productos` con cabecera `HX-Request: true`.
- Validaciones de negocio en servidor (precio negativo y nombre vacío devuelven HTTP 400 con `<hx-partial>` de error sin reemplazar la tabla).
- Eliminación sin recarga con `DELETE /productos/{id}` (HTTP 200 con cuerpo vacío) y error 404 para IDs inexistentes.

### Tests de base de datos (sqlc)

```bash
make test
```

Levanta una base limpia, genera `sqlc` y `templ`, corre `go test ./... -v` y borra los contenedores y volúmenes al terminar.

### Limpieza manual

```bash
make down
```

---

## Estructura del Proyecto

- `main.go`: servidor HTTP (`net/http`), rutas y handlers para vistas completas y fragmentos HTMX.
- `run-app.sh`: script con tareas previas, ejecución y limpieza posterior.
- `Makefile`: tareas `run-app`, `run`, `generate`, `ssr-test`, `test` y `down`.
- `views/`: componentes Templ (`.templ`) y archivos Go generados (`*_templ.go`):
  - `layout.templ`: plantilla HTML5 base con Pico.css, estilos locales y CDN de HTMX 4.0.
  - `index.templ`: página principal con contenedor `#mensajes-feedback`, formulario y listado.
  - `entity_form.templ`: formulario con `hx-post="/productos"`, `hx-target="#lista-productos"` y `hx-swap="outerHTML"`.
  - `entity_list.templ`: tabla de productos con botón de borrado (`hx-delete="/productos/{id}"`, `hx-target="closest tr"`, `hx-swap="outerHTML"`, `hx-confirm`).
  - `partials.templ`: fragmentos `<hx-partial>` para mensajes de feedback y limpieza del formulario.
- `logic/`: validaciones de reglas de negocio puras (`ValidateProduct`, `ValidateCategoria`, `ValidateCliente`).
- `db/schema/` y `db/queries/`: esquema SQL y consultas para `sqlc`.
- `db/sqlc/`: código generado y tests CRUD.
- `static/`: hojas de estilo `pico.min.css` y `styles.css`.
- `ssr_requests.hurl`: pruebas de integración de las rutas y fragmentos HTMX.
- `.devcontainer/`: `Dockerfile` y `docker-compose.yml`.
