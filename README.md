# Trabajo Especial - TP5: Renderizado del Lado del Servidor con Templ

Aplicación web de **MarkeTandil** para gestión de **productos** y **categorías**, implementando **Server-Side Rendering (SSR)** con **Templ**, estilos semánticos con **Pico.css** y el patrón **Post/Redirect/Get (PRG)** en Go.

---

## 🚀 Interfaz Web (TP5)

Al iniciar el servidor con `make run`, abrir en el navegador:

👉 **http://localhost:8080**

Permite visualizar la lista de productos de forma dinámica, dar de alta nuevos productos mediante el formulario clásico y eliminarlos, todo procesado directamente en el servidor Go con recargas de página completas.

---

## 📋 Requisitos

Solo hace falta tener instalado:

- **Docker** (con Docker Compose)
- **Make**

Go, `sqlc`, `templ`, PostgreSQL y Hurl se ejecutan dentro de los contenedores Docker, no hace falta instalarlos en el host.

---

## 🛠️ Comandos de Ejecución

### 1. Iniciar la aplicación

Desde la raíz del directorio `tpe`:

```bash
make run
```

Este comando:
1. Construye la imagen Docker y levanta los contenedores (PostgreSQL y Go).
2. Genera el código de acceso a datos con `sqlc generate`.
3. Compila los componentes de presentación con `templ generate`.
4. Inicia el servidor web en **http://localhost:8080**.

Cuando el servidor está listo, se muestra en consola:
```
Conectado a la base de datos correctamente
Servidor SSR con Templ escuchando en http://localhost:8080
```

### 2. Probar la integración SSR con Hurl

Con el servidor corriendo, en **otra terminal**:

```bash
make ssr-test
```

Ejecuta `ssr_requests.hurl`, validando el renderizado HTML inicial (`GET /`), el envío de formularios con redirección PRG (`POST /productos` con HTTP 303), y las validaciones de negocio en el servidor. Termina con:

```
ssr_requests.hurl: Success (5 request(s) in ...)
```

### 3. Tests unitarios y de base de datos (TP2)

```bash
make test
```

Levanta una base de datos limpia, ejecuta `sqlc generate`, compila vistas con `templ generate`, corre toda la suite de tests (`go test ./... -v`) y al finalizar detiene los contenedores.

### 4. Detener todo

Cortar el servidor con `Ctrl+C` y luego:

```bash
make down
```

Baja los contenedores y borra los volúmenes de la base de datos.

---

## 🏗️ Arquitectura y Componentes Templ

Siguiendo las filminas de la cátedra (`09-view-static.pdf`, `10-view-dynamic.pdf`, `07-logic.pdf`, `12-slides.pdf`):

### Directorio `views/`
- **`layout.templ`**: Estructura base del documento (`<!DOCTYPE html>`, `<head>`, `<body>`, encabezado) con enlace a **Pico.css** y estilos locales `/static/styles.css`. Aloja el contenido dinámico mediante `{ children... }`.
- **`entity_form.templ`**: Formulario tradicional (`<form method="POST" action="/productos">`) con campos para nombre, descripción, precio, stock y un `<select>` poblado dinámicamente con las categorías de la base de datos.
- **`entity_list.templ`**: Componente que recibe `[]db.Producto` y un mapa de categorías, renderizando una tabla semántica (`<table class="striped">`) con los datos formateados y formularios POST de eliminación por producto.
- **`index.templ`**: Componente `IndexPage` que ensambla el `Layout`, bloques de alerta para feedback (errores y éxitos) y organiza el formulario y el listado en la grilla visual.

### Flujo Post/Redirect/Get (PRG)
1. **Envío de formulario**: El usuario completa el formulario y envía `POST /productos` con `application/x-www-form-urlencoded`.
2. **Validación en tres capas**: El servidor Go parsea los campos y delega la validación de negocio en `logic.ValidateProduct`.
3. **Persistencia**: Si es válido, inserta con `s.queries.CrearProducto`.
4. **Redirección**: En lugar de renderizar directamente en el POST, devuelve `HTTP 303 See Other` redirigiendo a `/?success=...`.
5. **Carga limpia**: El navegador ejecuta `GET /`, evitando reenvíos accidentales si el usuario actualiza la página (F5).

---

## 📁 Estructura del Proyecto

```
main.go              servidor HTTP, rutas SSR y handlers del patrón PRG
views/               componentes templ (.templ) y código Go compilado (*_templ.go)
  layout.templ       plantilla base HTML5 con Pico.css
  entity_form.templ  componente formulario de alta
  entity_list.templ  componente tabla de productos y botón eliminar
  index.templ        componente de página principal (IndexPage)
logic/               capa de lógica pura del negocio (validaciones)
db/schema/           schema.sql con la definición de tablas
db/queries/          consultas SQL tipadas para sqlc
db/sqlc/             código generado por sqlc y tests unitarios
static/              archivos estáticos servidos en /static/ (pico.min.css, styles.css)
ssr_requests.hurl    pruebas de integración SSR y PRG con Hurl
DEFENSA_TP5.md       guía teórica y respuestas para la defensa ante la cátedra
.devcontainer/       Dockerfile (con sqlc, templ y hurl) y docker-compose.yml
Makefile             automatización de generación, ejecución y pruebas
```
